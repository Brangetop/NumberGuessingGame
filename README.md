Guess the Number (Go)

Simple command-line game written in Go to practice user input, random numbers, and basic control flow.
Overview

This program asks you to provide an upper bound N, then picks a random integer in the range 0..N (inclusive). You repeatedly guess numbers until you find the secret value; the program responds with "Too high!", "Too low!" or "You have guessed!".
Features

    Reads an upper bound from stdin
    Generates a random target in 0..N
    Reads guesses from stdin and gives immediate feedback
    Minimal, educational example for beginners

Requirements

    Go 1.20+ (or recent Go toolchain)

Usage

    Clone the repo.
    Run:

    go run .

    Follow prompts:
        Enter the maximum value (N)
        Enter guesses until you find the secret number

Notes / Improvements

    The repo is a learning exercise — not production-ready.
    Recommended improvements: seed the RNG with time, validate inputs and retry on invalid input, and provide attempt counts or a replay option. Maybe later.
