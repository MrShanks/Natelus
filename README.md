# The story

Once upon a time there was a great hero, called Natelus, with some strengths and weaknesses, as all
heroes have. After battling all kinds of monsters for more than a hundred years, Natelus now has the
following stats:
• Health : 70 - 100
• Strenght: 70 - 80
• Defence: 45 - 55
• Speed: 40 - 50
• Luck 10% - 30% ( 0% means no luck, 100% lucky all the time)
Also, he possesses 2 skills:
Rapid strike: Strike twice while it's his turn to attack, there's a 10% he'll use this skill every time he
attacks
Magic shield: takes only half of the usual damage when an enemy attacks, there's a 20% change he'll
use this skill every time he defends.


# Gameplay

As Natelus walks through a magical ever-green forest of the game, he encounters some wild beasts,
with the following properties:
• Health: 60 - 90
• Strenght: 60 - 90
• Defence: 40 - 60
• Speed: 40 - 60
• Luck: 25% - 40%
You'll have to simulate a battle between Natelus and a wild beast, either at command line or using a
web browser. On every battle, Natelus and the beast must be initialized with random properties,
within their ranges.
The first attack is done by the player with the higher speed. If both players have the same speed,
then the attack is carried on by the player with the highest luck. After an attack, the players switch
roles: the attacker now defends and the defender now attacks.
The damage done by the attacker is calculated with the following formula:
Damage = Attacker strength - Defender defence
The damage is subtracted from the defender's health. An attacker can miss their hit and do no
damage if the defender gets lucky that turn.
Natelus's skill occurs randomly, based on their chances, so take them into account on each turn.


# Game over

The game ends when one of the players remain without health or the number of turns reaches 20.
The application must output the results each turn:
• what happens?
• which skills were used (if any)?
• the damage done
• defender's health left
If we have a winner before the maximum number of rounds is reached, he must be declared.


# Rules

For this magical forest to be fair and work, you'll have to follow these rules:
• Write code in plain language of choice, preferably Python, Go or Java, whiteout any
frameworks (you are free to use 3rd parties for testing, or UI libs/frameworks)
• Make sure your application is decoupled, code reusable and scalable. For example, can a
new skill easily be added to our hero?
• Is your code bug-free and tested?