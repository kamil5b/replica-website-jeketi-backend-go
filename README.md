# STILL WORK IN PROGRESS. UNUSABLE UNTIL BETA VERSION

# replica-website-jeketi-backend-go
**Latest version: invdev**

A replica of jkt48.com backend, using Go.

# Project Roadmap
- [ ] Authentication
- [ ] Seeding active member prior to 2024
- [ ] User attachment system
- [ ] JKT48 Point system with no gateway payment
- [ ] OFC System
- [ ] CRUD Show 
- [ ] Show gacha system
- [ ] CRUD JKT48 Event (HS, VC, etc)
- [ ] Event war system
- [ ] CRUD Master data
- [ ] **INITIAL RELEASE v1.0.0**

Potential improvements after release:
- [ ] Send email
- [ ] OTP
- [ ] RBAC
- [ ] Suspension system
- [ ] User blacklist
- [ ] Gacha soft blacklist
- [ ] Gacha whitelist
- [ ] Gacha oshimen based
- [ ] MVPs
- [ ] Show discount based on age
- [ ] Concert attendance

# Folder Management
- root: use for system to run the app
- route: list api routes
- handler: use for validation and later process the service
- service: the business side of the service
- repository: communicate with database (manipulation and stuff)
- model: the base entities
- request: entities for requests
- constant: the constants of the app
- middleware: use for handling middlewares that later need in routes

# Entity Relationship Diagram
rough ERD, will be updated soon

[![ER Diagram](https://pbs.twimg.com/media/GRVF2lAaIAA2wjL?format=jpg)](https://x.com/kamil5b/status/1807427722031010060)
