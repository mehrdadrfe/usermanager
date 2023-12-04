# User Manager Package

The `usermanager` package provides functionality for managing users and their penalties in a simple penalty system.

## Overview

The `UserManager` type in this package is responsible for managing a collection of users, where each user has a unique identifier (`ID`), the number of penalty days (`PenaltyDays`), and the associated penalty amount (`PenaltyAmt`). Users can be added, penalties can be applied, and penalties can be checked and paid off.

## Usage

To use the `usermanager` package in your project, follow these steps:

1. **Install the Package:**

   ```bash
   go get github.com/mehrdadrfe/usermanager
