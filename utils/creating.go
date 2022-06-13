package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func CreateDir(path string) error {
	if (path != ".") && (path != "./") {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateFiles(path string, token string) error {
	file, err := os.Create(path + "/main.py")
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(`
import random, nextcord, os
from nextcord.ext import commands

description = """An example bot to showcase the nextcord.ext.commands extension
module.

There are a number of utility commands being showcased here."""

intents = nextcord.Intents.default()
intents.members = True

dotenv.load_dotenv()
TOKEN = os.getenv('TOKEN')
COGDIR = os.getenv('COGDIR') or './cogs'

class DaBot(commands.Bot):

	def __init__(self):
		super().__init__(command_prefix="$", description=description, intents=intents)

	async def on_ready():
		print(f"Logged in as {bot.user} (ID: {bot.user.id})")

bot = DaBot()

@bot.command()
async def add(ctx, left: int, right: int):
    """Adds two numbers together."""
    await ctx.send(left + right)


@bot.command()
async def roll(ctx, dice: str):
    """Rolls a dice in NdN format."""
    try:
        rolls, limit = map(int, dice.split("d"))
    except Exception:
        await ctx.send("Format has to be in NdN!")
        return

    result = ", ".join(str(random.randint(1, limit)) for r in range(rolls))
    await ctx.send(result)


@bot.command(description="For when you wanna settle the score some other way")
async def choose(ctx, *choices: str):
    """Chooses between multiple choices."""
    await ctx.send(random.choice(choices))


@bot.command()
async def repeat(ctx, times: int, content="repeating..."):
    """Repeats a message multiple times."""
    for i in range(times):
        await ctx.send(content)


@bot.command()
async def joined(ctx, member: nextcord.Member):
    """Says when a member joined."""
    await ctx.send(f"{member.name} joined in {member.joined_at}")


@bot.group()
async def cool(ctx):
    """Says if a user is cool.

    In reality this just checks if a subcommand is being invoked.
    """
    if ctx.invoked_subcommand is None:
        await ctx.send(f"No, {ctx.subcommand_passed} is not cool")


@cool.command(name="bot")
async def _bot(ctx):
    """Is the bot cool?"""
    await ctx.send("Yes, the bot is cool.")

bot.run(TOKEN)
	`))

	if err != nil {
		return err
	}

	file, err = os.Create(path + "/.env")
	if err != nil {
		return err
	}

	_, err = file.Write([]byte("TOKEN=" + token))

	if err != nil {
		return err
	}

	err = os.Mkdir(path+"/cogs", 0755)
	if err != nil {
		return err
	}

	file, err = os.OpenFile(path+"/cogs/example.py", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(`
from nextcord.ext import commands

class Cog(commands.Cog):

	def __init__(self, bot : commands.Bot):
		self.bot = bot

def setup(bot : commands.Bot):
	bot.add_cog(Cog(bot))	
	`))

	if err != nil {
		return err
	}

	file, err = os.OpenFile(path+"/requirements.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(`
git+https://github.com/nextcord/nextcord
python-dotenv
	`))

	if err != nil {
		return err
	}

	return nil

}

func InitializeGit(dir string) error {

	c := exec.Command("git", "init")
	c.Dir = dir

	err := c.Run()

	if err != nil {
		return err
	}

	c = exec.Command("git", "add", ".")
	c.Dir = dir

	err = c.Run()

	if err != nil {
		return err
	}

	c = exec.Command("git", "commit", "-m", "'Initial commit from create-nextcord-app'")
	c.Dir = dir

	err = c.Run()

	if err != nil {
		return err
	}

	return nil

}

func InitializeVenv(dir string) error {

	p := "python"
	pip := "./env/Scripts/pip3.exe"

	switch runtime.GOOS {
	case "windows":
		p = "python"
		pip = "./env/Scripts/pip3.exe"
	case "darwin":
		p = "python3"
		pip = "./env/bin/pip3"
	case "linux":
		p = "python3"
		pip = "./env/bin/pip3"
	}

	c := exec.Command(p, "-m", "venv", "env")
	c.Dir = dir

	err := c.Run()

	if err != nil {
		return err
	}

	c = exec.Command(pip, "install", "git+https://github.com/nextcord/nextcord", "python-dotenv")
	c.Dir = dir

	err = c.Run()

	if err != nil {
		return err
	}

	return nil

}
