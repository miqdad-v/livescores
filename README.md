# LiveScores

Live sports scores for Wear OS (Galaxy Watch), backed by a Go service.

## Planned architecture
- `backend/`: Go service that polls a sports API, caches results, and serves a REST API
- `watch/`: Kotlin + Jetpack Compose for Wear OS app
- `docs/`: design notes and diagrams

## Status
Early development.