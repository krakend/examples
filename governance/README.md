# Quotas example

Imagine we are game company with a catalog of games that share common mechanics,
and we want to have different limits to the resources to tweak game difficulty. 

We want single place to limit access to those resources.

Lets give some names to the games: 

- Galaga (`galaga`)
- Space Harrier (`space_harrier`)
- R-Type (`r_type`)

Every game has the following limited resources:

- **weapon_power_up**: This a simple power up, that will update a 
    backend service with the weapon power up consumed.

- **squad_call**: A request to the server will invoke a squad to help the
    main player. The **number of members that will come to help 
    depend on a backend call**, that will determine how many of this
    resource is consumed.

We also want to charge per credit (as if it was a virtual arcade), so
all games will have:

- **credits**: A call to credits consumes a credit for a each game. Credits
    can be consumed for any game.


For the players, we have two tiers:

- `freemium`
- `premium` 
  
The `premium` users have more **credits** to play each day / and hour (because
they are paying users). Since this is tied to our earnings, we are going to
user a **redis cluster** to limit this resource.

The tier for a player is extracted from its api key, using the roles feature,
so the request will use the `Authorization` header. For this example, we 
have two players : `bart` and `homer` (that we use as authorization keys).

In order to identify what game is using each player, we user an additional
header `X-Game` (that can have one of these values: `galaga`, `space_harrier`
or `r_type`).

## Usage

In order to run, we need to spawn a redis cluster and a standalone redis instance.
There is a script under the `./redis` dir, called `start.sh` that spawns 
redis instances with host network, and then executes a command to create a cluster.

Once redis is in place, place your `LICENSE` in the current dir and run:

```
krakend run -c krakend.json
```

In the `clients` dir there is a `curl.sh` script, that you can edit to
tweak the variables before making a request:

- NUM_REQUESTS: number of sequential requests to make
- QUOTAS_ENDPOINT: endpoint to call
- QUOTAS_USER: the player to use
- QUOTAS_GAME: the game to play
