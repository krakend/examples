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


## Mission Control: quota events dashboard

> **Note:** quota events are available since KrakenD Enterprise **3.1**.

Every quota in this example also emits **events** that end up in a Grafana dashboard:

```
KrakenD quotas ──(governance/event-recorders, "mem")──► async_agent ──► Elasticsearch ──► Grafana
```

- `governance/event-recorders` declares an in-memory recorder (`game_events`), and every quota
    processor points to it with `"event_recorder": "game_events"`.
- The quotas emit `quota.limit.increment` (every time a resource is consumed) and
    `quota.limit.reached` (when a pilot exhausts a resource). Rejected requests are not emitted.
- The `quota-events-to-elasticsearch` async agent consumes the recorder topic and indexes
    each event into the `quota-events` data stream.
- The `elasticsearch-setup` container installs an ingest pipeline (`config/elasticsearch`) that
    flattens the events and adds some game context: the `ship`, the player `plan`, and when the
    pilot can play again (`reset_at`, computed from the event's `nextRefill`).

For production you would replace the `mem` recorder with a `kafka` one, and run the async agent
in a separate KrakenD.

Start everything and play some games:

```
docker compose up -d
./client/play.sh 50        # or ./client/play.sh 0 to play until Ctrl+C
```

Then open [http://localhost:3000](http://localhost:3000) (no login needed) to see **Earth vs Aliens: Mission Control**:

- **Coins inserted, pilots in action, power-ups fired, bombs delivered, GAME OVER** counters.
- **Battle activity**: resources consumed over time.
- **Coins by plan**: credits per `premium` / `freemium`. The `credits` quota uses `"hash_keys": true`,
    so pilots are anonymous in that quota's events.
- **High scores**: pilots ranked by power-ups and bombs.
- **Hangar**: which ships burn more resources.
- **Weapon heat / Bomb bay load**: how much of the current hourly window each pilot and ship has used.
    The bomb bay can go over 100% because the weight comes from the backend response, so the last
    reload is accounted after it is served.
- **GAME OVER feed**: who ran out of what, and when they can continue.

Each limit (`hour`, `day`) emits its own events, so the **Window** selector picks the limit to count.

> Keep in mind that the quota key in this example is the API key itself (the `Authorization` header).
> Without `hash_keys` it is stored in clear in the events. In a real setup use a user id (e.g. from a
> JWT claim) or enable `hash_keys`.
