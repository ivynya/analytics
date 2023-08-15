
# Changelog

### Analytics V3
1. Service now written in Go for resource efficiency and speed.
2. New `GET` `/v3/campaign/:id/interaction/:iID` endpoint to get campaign data for a specific interaction
3. JSON responses are now returned in streamlined Page schema
4. Deprecated V1 API - will continue to function but will be removed in a future release (see Migration Guide)

### Analytics V2
1. New endpoints and database schema to handle Interactions, a granular way to track campaign performance
2. New endpoint to get campaign data without creating a visit

### Analytics V1
1. Initial release