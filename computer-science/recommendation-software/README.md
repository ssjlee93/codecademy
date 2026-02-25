# Portfolio Project: Recommendation Software

## Think of an idea
Pokemon recommendation software.  

use [PokeAPI](https://pokeapi.co) to get data.  

**What kind of recommendation system do you want to build?**  
Pokemon recommendation  

**Who do you want to build something for? Yourself? Your friends? Your family? Your co-workers?**  
Pokemon players  

**What topics will be searchable in your recommendation system?**  
Pokemon types - keep it minimal for now

### MVP
recommend pokemon randomly.  
recommend pokemon based on type and recommend randomly.  

### future ideas

#### recommend party for a given game
when i'm playing a game, i want to know the least number of pokemons i can carry for all my HMs.  
when i'm playing a game, I want to optimize my party for the gym leaders - requires more data than pokemoin.  

## Project brainstorming

**What does your program do?**  
Recommends a random pokemon based on the game I'm playing.  
Mainly HM and field move pokemons.  
Or which pokemon to catch in which version of the game.  

**What data do you need?**  
Pokemon data  
- Pokemon sprite variations.  
- Which Pokemon appears in each game.  
- The latest game in which the Pokemon is available.  
- Learnable HMs

**What questions will you ask the user?**  
The type of Pokemon they want a recommendation for.  


**How do the above questions return a recommendation?**  
I just randomly pick out a pokemon of the type.  

## ideas for what the application does  

I can split the application to do 2 different things :  
- Crawl data on all pokemon  
- Serve data on all collected pokemons

I would need 2 cmds to crawl data on one side and expose REST APIs on the data I need.  

