# Road map for next release

1. State support (state pattern)

  > _The bot should be able to react differently to it's environment depending on a given state. An example of this can be the state `pause` where the bot does not run hooks, commands, services, etc. until the bot has been activated again. This must use a key=value database in case of restarts. An example: The Guild has a bot to assign roles to every new member. This role is called `member` and lets users write messages in the `#general` channel. But spammers keep joining the guild, so the moderator/admin/owner might be interested in changing the state of the bot so it gives a different role to new members. The owner writes `!state lockdown` and the bot sets it's state to `lockdown` (a custom state). The bot no longer assign the `member` role to new members, but rather `suspect`. The owner can then easily review users by filtering using their distinct role and evaluate each one. When the threat is over, the bot can be set to normal again by `!state normal`. Members will now gain the `member` role._

2. The owner has absolute power

  > _The guild owner should always have complete bot access. This is so after the bot has been added, the owner should be able to add permissions for certain roles or specific members using commands. Reducing the need for hard coding any snowflake tokens or weak bot behavior._

3. Commands needs flags and a third lib pattern

  > _Adding flags with proper parsing increases the flexibility and power of the bot, but doing this correctly can be difficult and thinking ahead might prove to be even a bigger challenge. Therefore using a framework to handle program arguments, as if each command was a software, one can take advantage of libraries such as `urfave/cli` which has implementing flags with aliasing and documentation of usecases. This leaves less work to us, and we can focus on improving the library in stead._

4. ~~Environmental variables~~

  > _The cloud is becoming vastly popular and using solutions such as docker and Heroku seems like a smart way to run discord bots. With both of these solutions, configurations containing sensitive information is done using environment variables. Since these can easily be exported in any linux environment, I believe the Unison framwork should look for three different environment variables, UNISON_DISCORD_TOKEN, UNISON_COMMAND_PREFIX, UNISON_STATE. Where UNISON_DISCORD_TOKEN is required, but can be set using the Settings struct as an alternative. UNISON_COMMAND_PREFIX defaults to triggering commands using mention if not specified. UNISON_STATE is the default bot state, as discussed in 1\. point. This will default to STATE=normal{0}. There should be preserved states for the Unison framework.._

5. ~~Reserved bot state~~

  > _States are indexed within the range of [0, 255] where as states[0, 10] are preserved. These cannot be set by a bot. state{0}=normal, state{1}=pause, state{2}=debug_

6. New parameter for ~~commands~~ and hooks.

  > _A potential eternal loop, is when a bot reacts to new messages with a reply. What if that message was written by the bot itself? The hook should be aware if it's from itself, but not be forced by Unison to not run on bot-related events. This is useful in case messages needs to be logged._

7. Thread hooks and commands

  > _Commands and hooks should not be blockers to prevent Unison from handling incoming requests. Should add channels for proper chaining of commands._
