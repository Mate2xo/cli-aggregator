package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Mate2xo/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("please enter one argument as the username")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("could not find user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set current user as: %w", err)
	}

	fmt.Printf("Current user has been set to %s.", user.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("please enter one argument as the username")
	}

	username := cmd.args[0]
	if username == "" {
		return errors.New("please enter a username")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      username,
	})
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}

	fmt.Printf("User %s was created\n", user.Name)
	fmt.Printf("User: %+v\n", user)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not list users: %w", err)
	}

	for _, user := range users {
		msg := fmt.Sprintf("* %s", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			msg += " (current)"
		}
		println(msg)
	}

	return nil
}
