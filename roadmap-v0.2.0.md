# Road map for next release

1. State support (state pattern)

⋅⋅⋅_The bot should be able to react differently to it's environment depending on a given state. An example of this can be the state `pause` where the bot does not run hooks, commands, services, etc. until the bot has been activated again. An example: The Guild has a bot to assign roles to every new member. This role is called `member` and lets users write messages in the `#general` channel. But spammers keep joining the guild, so the moderator/admin/owner might be interested in changing the state of the bot so it gives a different role to new members. The owner writes `!state lockdown` and the bot sets it's state to `lockdown` (a custom state). The bot no longer assign the `member` role to new members, but rather `suspect`. The owner can then easily review users by filtering using their distinct role and evaluate each one. When the threat is over, the bot can be set to normal again by `!state normal`. Members will now gain the `member` role._

1. test
