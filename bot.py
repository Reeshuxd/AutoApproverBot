from os import getenv
from asyncio import sleep
from pyrogram import Client, filters
from pyrogram.errors import FloodWait
from logging import basicConfig, INFO
from pyrogram.types import InlineKeyboardButton, InlineKeyboardMarkup

basicConfig(level=INFO)
client = Client(
    "ApproverBot",
    api_hash="eb06d4abfb49dc3eeb1aeb98ae0f581e",
    api_id=6,
    bot_token=getenv("TOKEN")
)

@client.on_message(filters.command("start"))
async def start(c, m):
    text = """
<b>Hello <a href="tg://user?id={}">{}</a></b>
I am a bot made for accepting upcoming join requests at the time they comes.
I am a bot made using <a href="python.org">python</a> for a better performance.

<i>Bot made with 💝 by <a href="t.me/aboutreeshu">reeshu</a> for you!</i>
<b>Support Chat:</b> @UserChatRoom
    """
    await m.reply(
        text.format(m.from_user.id, m.from_user.first_name), 
        parse_mode="html",
        disable_web_page_preview=True,
        reply_markup=InlineKeyboardMarkup(
            [[InlineKeyboardButton("My Source Code", url="github.com/reeshuxd/AutoApproverBot")]]
        )
    )

@client.on_chat_join_request(filters.channel)
async def jn(c, m):
    try:
        await client.approve_chat_join_request(m.chat.id, m.from_user.id)
    except FloodWait as fd:
        await sleep(fd.x + 2)
    except BaseException:
        pass

client.run()