# Meeting Room Booking System

## Overview

The Meeting Room Booking System is a robust application designed to facilitate the booking of meeting rooms within an organization. The system allows users to reserve rooms for specific durations while ensuring thread-safe operations and handling concurrency.

## Features

- **Room Management:** Add and retrieve meeting rooms.
- **User Management:** Manage users who can book meeting rooms.
- **Booking System:** Create, reschedule, and check booking status.
- **Concurrency Handling:** Safe operations to prevent race conditions.
- **Error Handling:** Graceful management of errors and room availability checks.
- **Singleton Pattern:** Ensures consistent state management across the application.

## Architecture

The application is organized into the following components:

- **Models:**
  - `Room`: Represents a meeting room with attributes like `RoomID`, `Name`, and `Capacity`.
  - `User`: Represents a user with attributes like `UserID`, `Name`, and `Email`.
  - `Booking`: Represents a booking with attributes like `BookingID`, `Room`, `User`, `StartTime`, `EndTime`, and `Status`.

- **Services:**
  - `RoomManager`: Singleton responsible for managing rooms.
  - `BookingService`: Singleton responsible for managing bookings.

## UML Diagram

![UML Diagram](path/to/your/UML_diagram.png)

## Getting Started

### Prerequisites

- Go 1.16 or later
- A working Go environment

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/meeting-room-booking-system.git
   cd meeting-room-booking-system
