# Quotas example

Imagine we are game company with a shot'em up game called **Earth vs Aliens**,
and we want to have different limits for the game resources.

In **Earth vs Aliens** the player can choose one of this ships to play:

- Ala-X (`ala-x`): a ship with fast fire weapons with several power ups,
    but that has only capacity to transport a **single bomb**.
- Ala-B (`ala-b`): a balanced ship, that can transport up to **two bombs**,
    and has a fer power ups for its fast weapons.
- Ala-Y (`ala-y`): a bombardier ship that has capacity for **4 bombs**
    but with less main weapon power ups.

So we define the following limited resources:

- **weapon_power_up**: This a simple power up, that will update a 
    backend service with the weapon power up consumed.

- **bomb reload**: A support ship will come and reload the empty
    bomb slots of the player's ship: A request to the server will 
    the check the empty slots in the ship, and will tell how many
    bombs the support ship must provide.

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
header `X-Ship` (that can have one of these values: `ala-x`, `ala-b`
or `ala-y`).

## The Config file

In the `./config/krakend/krakend.json` file we can take a look at some
details:

- we set up `auth/api-keys` to have authorized players, each one with 
    its own "pay tier"
- we have the **two kind of redis connections**: "client pool" / "cluster":
    the cluster one is set to track the credits (so if we had additional
    games we could track them in the same place, and with more capacity)
- we have `quotas` for each of the resources, and different rules 
    (in our case are like the "levels" for each resource).
- for each endpoint then we have the quota that it must use, and depending
    on the tier, what rule inside that quota to apply.
- in order to identify a unique user, we directly use the `Authorization` header
    but a user id could be extracted from a JWT claim.
- in endpoint `/request_bomb_reload` we put a backend that serves static
    files from the `./confing/krakend/fake_backend/data/` directory: 
    by changing the `consume.json` file there, we can change how many
    **tokens to consume** are in the response.

## Usage

Place your `LICENSE` in the `./config/krakend` directory an execute:

```
docker compose up -d
```

In the `clients` dir there is a `curl.sh` script, that you can edit to
tweak the variables before making a request:

- NUM_REQUESTS: number of sequential requests to make
- QUOTAS_ENDPOINT: endpoint to call
- QUOTAS_USER: the player to use
- QUOTAS_SHIP: the game to play

