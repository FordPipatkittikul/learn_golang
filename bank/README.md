# Hexagonal architecture

Ports & Adapters architectural pattern

![alt text](<Screenshot 2026-01-14 204003.png>)

**handler**

Handles HTTP requests and responses.

→ Parse input, call service, return HTTP result.

**service**

Contains business logic.

→ Defines service interfaces and their implementations.

→ Uses repositories to get/save data.

→ Uses errs for business-level errors.

**repository**

Handles data access.

→ Defines repository interfaces and implementations (DB, external APIs).

→ No business logic here.

**mock**

Fake implementations for testing.

→ Used in unit tests to avoid real DB or services.

**logs**

Centralized logging configuration and helpers.

→ Ensures consistent logging across the app.

**errs**

Shared error definitions.

→ Standardized errors (e.g. NotFound, ValidationError, InternalError).

**Quick Note**

Interface = what a component can do

Implementation = how it does it. Using structand New func

## Layer of Architecture

Presentation -> Business -> Database

Presentation : handler

Business : service layer (this is where we should handle log)

Database : repository layer