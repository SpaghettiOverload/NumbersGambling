# Guess The Number Game
**/ WITH A PINCH OF A GAMBLING /**

*The sole purpose of creating this simple game is practicing Go as a second language that comes in my stack after Python.* 
<br>
###Game mechanics
The game starts with a player. 
###### As a player you get:
    Main balance: $100
    JACKASS credit: 1
    HINTS: 3

At the beginning of each round, a base number is shown, then a new unknown one is generated. You have to place a bet about the new number.
######Bet variants:
    A) Jackpot guess. 
    *directly trying to guess the value of the new number.
    
    B) Movement prediction.
    *a guess whether the new number value is less or greater than the base number.

Keep in mind that bets come at a price, though.
At the end of a round, the unknown number become a base number for the next round. 

##Game rules
Conditions

Winning | Losing
------------ | -------------
$1,000,000 | Balance < 0


Pricing

Item | Cost
------------ | -------------
Jackpot guess | JACKASS credit + $50
JACKASS credit | $2500
Incorrect prediction | $50  
Winning strike failure | 50% from balance + 50$
Range reduction* | $50,000
HINT** | FREE

* *For $50,000 you can purchase a "Number range reduction" with 10, meaning that next round, the new number to guess will be generated in the range of 0 to (100 minus (n x 10). Range cannot be less than 10!    
* **3 hints are available in total and 1 hint can be used per round. Hinting is disabled while in Wining Strike State.

1. A correct "Jackpot guess" grants you $1,000,000 and makes you an instant winner.
   - You can do Jackpot guesses as many times as you want per round, as long as you have JACKASS credits.  
2. A correct "Movement prediction" will put you into a temporary "winning strike" state. At this point the only valid commands will be "U", "D" and "C" (*See bottom of the document for info)
######Winning strike state
    The first correct guess will grant you a temporary winning balance of $200.
    - From here you can:
        a) Cash-out to main balance and start next round.
        b) Continue guessing:
            # Every next successful guess will quadruple your temporary balance.
            # The first wrong guess will make you lose your temporary balance and fine you 50% from your main balance + $50.



###Valid player commands*
* "B" - *showing your balance state.*
* "H" - *requesting a HINT for the number change.*
* "U" - *a guess that the new number went UP.*
* "D" - *a guess that the new number went DOWN.*
* "J" - *a wish to purchase 1 JACKASS credit.*
* "R" - *a wish to purchase number range reduction.*
* INT - *whole number for a Jackpot guess.*
* "C" - *cash out. Active while in winning strike state.*  

*Commands are case-insensitive.
