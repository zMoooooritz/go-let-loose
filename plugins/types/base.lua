---@meta

---Player information
---@class PlayerInfo
---@field Name string The player's display name
---@field ID string The player's unique identifier (Steam ID, etc.)
local PlayerInfo = {}

---Detailed player information with additional stats and status
---@class DetailedPlayerInfo : PlayerInfo
---@field Name string The player's display name
---@field ID string The player's unique identifier
---@field ClanTag string The player's clan tag
---@field Platform string Player's platform ("steam", "epic", "xbl", "none")
---@field Team string The player's team ("Allies", "Axis", "None")
---@field Faction string The player's faction ("US", "GER", "SOV", etc.)
---@field Role string The player's role ("Rifleman", "Officer", etc.)
---@field Unit Unit The player's unit/squad information
---@field Loadout string The player's current ladout
---@field Kills number Player's kill count
---@field Deaths number Player's death count
---@field Score Score Player's score breakdown
---@field Level number Player's level
---@field Position Position Player's world position
local DetailedPlayerInfo = {}

---Weapon information  
---@class Weapon
---@field ID string Weapon identifier (e.g., "M1_GARAND")
---@field Name string Human-readable weapon name (e.g., "M1 Garand")
---@field Category string Weapon category (Rifle, SMG, etc.)
---@field Factions table Array of factions that can use this weapon
local Weapon = {}

---Unit/Squad information
---@class Unit
---@field Name string Unit name (e.g., "Able", "Baker", "Command")
---@field ID number Unit ID
local Unit = {}

---Player score breakdown
---@class Score
---@field Combat number Combat score
---@field Offense number Offensive score
---@field Defense number Defensive score
---@field Support number Support score
local Score = {}

---3D position in game world
---@class Position
---@field X number X coordinate
---@field Y number Y coordinate
---@field Z number Z coordinate (elevation)
local Position = {}

---Team score data
---@class TeamData
---@field Allies number Allied team score/count
---@field Axis number Axis team score/count
local TeamData = {}

---Game map information
---@class GameMap
---@field ID string Map identifier
---@field Name string Full map name
---@field Tag string Short map tag
---@field PrettyName string Display name
---@field ShortName string Abbreviated name
---@field Allies string Allied faction
---@field Axis string Axis faction
---@field Orientation string Map orientation
local GameMap = {}

---Layer information (map + game mode)
---@class Layer
---@field ID string Layer identifier
---@field GameMap GameMap Map information
---@field GameMode string Game mode
---@field Attackers string Attacking team
---@field Environment string Environmental conditions
local Layer = {}

---Current game state
---@class GameState
---@field PlayerCount TeamData Player count per team
---@field GameScore TeamData Current scores
---@field RemainingSeconds number Time remaining in match
---@field CurrentMap Layer Current map/layer
---@field NextMap Layer Next map/layer
local GameState = {}

---Admin information
---@class Admin
---@field Name string Admin's name
---@field ID string Admin's player ID
---@field Role string Admin role ("owner", "senior", "junior", "spectator")
---@field Comment string Admin comment/description
local Admin = {}

---Server ban information
---@class ServerBan
---@field Type string Ban type ("temp" or "permanent")
---@field Player PlayerInfo Banned player info
---@field Timestamp string Ban timestamp
---@field Duration number Ban duration (for temp bans)
---@field Reason string Ban reason
---@field AdminName string Admin who issued the ban
---@field RawLog string Raw log entry
local ServerBan = {}

---Server configuration
---@class ServerConfig
---@field Name string Server name
---@field BuildNumber string Build number
---@field BuildRevision string Build revision
---@field SupportedPlatforms table Array of supported platforms
---@field PasswordProtected boolean Whether server is password protected
local ServerConfig = {}

---Session information
---@class SessionInfo
---@field ServerName string Server name
---@field MapName string Current map name
---@field GameMode string Current game mode
---@field RemainingMatchTime number Time remaining in match
---@field MatchTime number Total match time
---@field AlliedFaction number Allied faction ID
---@field AxisFaction number Axis faction ID
---@field MaxPlayerCount number Maximum players
---@field AlliedScore number Allied team score
---@field AxisScore number Axis team score
---@field PlayerCount number Current player count
---@field AlliedPlayerCount number Allied player count
---@field AxisPlayerCount number Axis player count
---@field MaxQueueCount number Maximum queue size
---@field QueueCount number Current queue size
---@field MaxVipQueueCount number Maximum VIP queue size
---@field VipQueueCount number Current VIP queue size
local SessionInfo = {}

---Command information
---@class Command
---@field ID string Command identifier
---@field Name string Command name
---@field ClientSupported boolean Whether command is client-supported
local Command = {}

---Command details with parameters
---@class CommandDetails
---@field Name string Command name
---@field Text string Command text
---@field Description string Command description
---@field DialogueParameters table Array of DialogueParameter
local CommandDetails = {}

---Dialogue parameter for commands
---@class DialogueParameter
---@field Type string Parameter type
---@field Name string Parameter name
---@field ID string Parameter ID
---@field DisplayMember string Display member
---@field ValueMember string Value member
local DialogueParameter = {}

---Server view with team and squad information
---@class ServerView
---@field Allies TeamView Allied team view
---@field Axis TeamView Axis team view
---@field Neutral SquadView Neutral players
local ServerView = {}

---Team view with commander and squads
---@class TeamView
---@field Commander DetailedPlayerInfo Team commander
---@field Squads table Map of squad name to SquadView
local TeamView = {}

---Squad view with players
---@class SquadView
---@field Team string Squad's team
---@field SquadType string Squad type ("Infantry", "Recon", "Armor")
---@field Name string Squad name
---@field Players table Array of DetailedPlayerInfo
local SquadView = {}

---Log entry
---@class LogEntry
---@field Timestamp string Log timestamp
---@field Level string Log level
---@field Message string Log message
---@field Raw string Raw log line
local LogEntry = {}