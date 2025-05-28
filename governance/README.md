# Quotas example

Lets imagine we have several games, that share common mechanics, so we
have a single place to limit access to resources.

Lets say those games are: 

- Galaga (`galaga`)
- Space Harrier (`space_harrier`)
- R-Type (`r_type`)

Every game has the following limited resources:

- **weapon_power_up**: This a simple power up, that will update a 
    backend service with the powerup.

- **squad_call**: A request to the server will invoke a squad to help the
    main player. The **number of members that will come to help 
    depend on a backend call**, that will determine how many of this
    resource is consumed.

- **lives**: A call to lives consumes adds a new live to the game


For the players, we have two tiers:

- `freemium`
- `premium` 
  
The `premium` users have more **lives** to play each day / and hour (because
are paying users). Since this is tied to our earnings, we are going to
user a **redis cluster** to limit this resource.

The tier for a player is extracted from its api key, using the roles feature,
so the request will use the `Authorization` header. For this example, we 
have two players : `bart` and `homer` (that we use as authorization keys).

In order to identify what game is using each player, we user an additional
header `X-Game` (that can have one of these values: `galaga`, `space_harrier`
or `r_type`).
