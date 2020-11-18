package blizzard_api

import (
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
)

type WoWClient struct {
	ApiClient
}

func NewWoWClient(region Region, redisClient *redis.Client, validCacheStatus []int, classic bool) *WoWClient {
	if validCacheStatus == nil {
		validCacheStatus = []int{200, 404}
	}

	return &WoWClient{
		ApiClient{
			httpClient:       new(http.Client),
			redisClient:      redisClient,
			game:             "wow",
			region:           region,
			validCacheStatus: validCacheStatus,
			Classic: classic,
		},
	}
}

// Achievement API

func (client *WoWClient) AchievementCategoryIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/achievement-category/index", options)
}

func (client *WoWClient) AchievementCategory(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/achievement-category/%d", options, id)
}

func (client *WoWClient) AchievementIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/achievement/index", options)
}

func (client *WoWClient) Achievement(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/achievement/%d", options, id)
}

func (client *WoWClient) AchievementMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/achievement/%d", options, id)
}

// Auction

func (client *WoWClient) Auction(connectedRealmId int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/connected-realm/%d/auctions", options, connectedRealmId)
}

// Azerite

func (client *WoWClient) AzeriteEssenceIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/azerite-essence/index", options)
}

func (client *WoWClient) AzeriteEssence(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/azerite-essence/%d", options, id)
}

func (client *WoWClient) AzeriteEssenceMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/azerite-essence/%d", options, id)
}

// Connected Realm

func (client *WoWClient) ConnectedRealmIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/connected-realm/index", options)
}

func (client *WoWClient) ConnectedRealm(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/connected-realm/%d", options, id)
}

// Covenant

func (client *WoWClient) CovenantIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/covenant/index", options, id)
}

func (client *WoWClient) Covenant(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/covenant/%d", options, id)
}

func (client *WoWClient) CovenantMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, DYNAMIC_NS,  "/covenant/%d", options, id)
}

func (client *WoWClient) SoulbindIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/soulbind/index", options, id)
}

func (client *WoWClient) Soulbind(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/soulbind/%d", options, id)
}

func (client *WoWClient) ConduitIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/conduit/index", options, id)
}

func (client *WoWClient) Conduit(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/conduit/%d", options, id)
}

// Creatures

func (client *WoWClient) CreatureFamilyIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/creature-family/index", options)
}

func (client *WoWClient) CreatureFamily(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/creature-family/%d", options, id)
}

func (client *WoWClient) CreatureTypeIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/creature-type/index", options)
}

func (client *WoWClient) CreatureType(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/creature-type/%d", options, id)
}

func (client *WoWClient) Creature(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/creature/%d", options, id)
}

func (client *WoWClient) CreatureDisplayMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/creature-display/%d", options, id)
}

func (client *WoWClient) CreatureFamilyMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/creature-family/%d", options, id)
}

// Guild Crest

func (client *WoWClient) GuildCrestComponentsIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/guild-crest/index", options)
}

func (client *WoWClient) GuildCrestBorderMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/guild-crest/border/%d", options, id)
}

func (client *WoWClient) GuildCrestEmblemMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/guild-crest/emblem/%d", options, id)
}

// Item

func (client *WoWClient) ItemClassIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/item-class/index", options)
}

func (client *WoWClient) ItemClass(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/item-class/%d", options, id)
}

func (client *WoWClient) ItemSetIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/item-set/index", options)
}

func (client *WoWClient) ItemSet(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/item-set/%d", options, id)
}

func (client *WoWClient) ItemSubclass(classId int, id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/item-class/%d/item-subclass/%d", options, classId, id)
}

func (client *WoWClient) Item(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/item/%d", options, id)
}

func (client *WoWClient) ItemMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/item/%d", options, id)
}

func (client *WoWClient) ItemSearch(options *RequestOptions, searchQuery string) *ApiResponse {
	return client.ApiRequest(SEARCH, STATIC_NS, "/item", options, searchQuery)
}

// Journal

func (client *WoWClient) JournalExpansionIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/journal-expansion/index", options)
}

func (client *WoWClient) JournalExpansion(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/journal-expansion/%d", options, id)
}

func (client *WoWClient) JournalEncounterIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/journal-encounter/index", options)
}

func (client *WoWClient) JournalEncounter(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/journal-encounter/%d", options, id)
}

func (client *WoWClient) JournalInstanceIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/journal-instance/index", options)
}

func (client *WoWClient) JournalInstance(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/journal-instance/%d", options, id)
}

func (client *WoWClient) JournalInstanceMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/journal-instance/%d", options, id)
}

// Mount

func (client *WoWClient) MountIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/mount/index", options)
}

func (client *WoWClient) Mount(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/mount/%d", options, id)
}

// Mythic Keystone Dungeon API

func (client *WoWClient) MythicKeystoneDungeonIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/mythic-keystone/dungeon/index", options)
}

func (client *WoWClient) MythicKeystoneDungeon(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/mythic-keystone/dungeon/%d", options, id)
}

func (client *WoWClient) MythicKeystoneIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/mythic-keystone/index", options)
}

func (client *WoWClient) MythicKeystonePeriodIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/mythic-keystone/period/index", options)
}

func (client *WoWClient) MythicKeystonePeriod(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/mythic-keystone/period/%d", options, id)
}

func (client *WoWClient) MythicKeystoneSeasonIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/mythic-keystone/season/index", options)
}

func (client *WoWClient) MythicKeystoneSeason(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/mythic-keystone/season/%d", options, id)
}

// Mythic Keystone Leaderboards Index

func (client *WoWClient) MythicKeystoneLeaderboardIndex(connectedRealmId int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/connected-realm/%d/mythic-leaderboard/index", options, connectedRealmId)
}

func (client *WoWClient) MythicKeystoneLeaderboard(connectedRealmId int, dungeonId int, periodId int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/connected-realm/%d/mythic-leaderboard/%d/period/%d", options, connectedRealmId, dungeonId, periodId)
}

// Mythic Raid Leaderboard API

func (client *WoWClient) MythicRaidLeaderboard(raid string, faction string, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/leaderboard/hall-of-fame/%s/%s", options, raid, faction)
}

// Pet

func (client *WoWClient) PetIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/pet/index", options)
}

func (client *WoWClient) Pet(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/pet/%d", options, id)
}

// Playable Class

func (client *WoWClient) PlayableClassIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/playable-class/index", options)
}

func (client *WoWClient) PlayableClass(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/playable-class/%d", options, id)
}

func (client *WoWClient) PlayableClassMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/playable-class/%d", options, id)
}

func (client *WoWClient) PlayableClassPvPTalentSlots(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/playable-class/%d/pvp-talent-slots", options, id)
}

// Playable Race

func (client *WoWClient) PlayableRaceIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/playable-race/index", options)
}

func (client *WoWClient) PlayableRace(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/playable-race/%d", options, id)
}

// Playable Specialization

func (client *WoWClient) PlayableSpecializationIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/playable-specialization/index", options)
}

func (client *WoWClient) PlayableSpecialization(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/playable-specialization/%d", options, id)
}

func (client *WoWClient) PlayableSpecializationMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/playable-specialization/%d", options, id)
}

// Power type

func (client *WoWClient) PowerTypeIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/power-type/index", options)
}

func (client *WoWClient) PowerType(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/power-type/%d", options, id)
}

// Profession

func (client *WoWClient) ProfessionIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/profession/index", options)
}

func (client *WoWClient) Profession(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/profession/%d", options, id)
}

func (client *WoWClient) ProfessionSkillTier(professionId int, skillTierId int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/profession/%d/skill-tier/%d", options, professionId, skillTierId)
}

func (client *WoWClient) ProfessionMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/profession/%d", options, id)
}

func (client *WoWClient) Recipe(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/recipe/%d", options, id)
}

func (client *WoWClient) RecipeMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/recipe/%d", options, id)
}

// PvP Season

func (client *WoWClient) PvPSeasonIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/pvp-season/index", options)
}

func (client *WoWClient) PvPSeason(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/pvp-season/index/%d", options, id)
}

func (client *WoWClient) PvPLeaderBoardIndex(pvpSeasonId int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/pvp-season/%d/pvp-leaderboard/index", options, pvpSeasonId)
}

func (client *WoWClient) PvPLeaderBoard(pvpSeasonId int, bracket string, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/pvp-season/%d/pvp-leaderboard/%s", options, pvpSeasonId, bracket)
}

func (client *WoWClient) PvPRewardsIndex(pvpSeasonId int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/pvp-season/%d/pvp-reward/index", options, pvpSeasonId)
}

// PvPTier

func (client *WoWClient) PvPTierMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/pvp-tier/%d", options, id)
}

func (client *WoWClient) PvPTierIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/pvp-tier/index", options)
}

func (client *WoWClient) PvPTier(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/pvp-tier/%d", options, id)
}

// Quest

func (client *WoWClient) QuestIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/quest/index", options)
}

func (client *WoWClient) Quest(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/quest/%d", options, id)
}

func (client *WoWClient) QuestCategoryIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/quest/category/index", options)
}

func (client *WoWClient) QuestCategory(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/quest/category/%d", options, id)
}

func (client *WoWClient) QuestAreaIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/quest/area/index", options)
}

func (client *WoWClient) QuestArea(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/quest/area/%d", options, id)
}

func (client *WoWClient) QuestTypeIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/quest/type/index", options)
}

func (client *WoWClient) QuestType(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/quest/type/%d", options, id)
}

// Realm

func (client *WoWClient) RealmIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/realm/index", options)
}

func (client *WoWClient) Realm(realmSlug string, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/realm/%s", options, realmSlug)
}

func (client *WoWClient) RealmById(realmId int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/realm/%d", options, realmId)
}

// Region

func (client *WoWClient) RegionIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/region/index", options)
}

func (client *WoWClient) Region(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/region/%d", options, id)
}

// Reputation

func (client *WoWClient) ReputationFactionIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/reputation-faction/index", options)
}

func (client *WoWClient) ReputationFaction(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/reputation-faction/%d", options, id)
}

func (client *WoWClient) ReputationTierIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/reputation-tiers/index", options)
}

func (client *WoWClient) ReputationTier(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/reputation-tiers/%d", options, id)
}

// Spell API

func (client *WoWClient) Spell(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/spell/%d", options, id)
}

func (client *WoWClient) SpellMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/spell/%d", options, id)
}

func (client *WoWClient) SpellSearch(options *RequestOptions, searchQuery string) *ApiResponse {
	return client.ApiRequest(SEARCH, STATIC_NS, "/spell", options, searchQuery)
}

// Talent API

func (client *WoWClient) TalentIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/talent/index", options)
}

func (client *WoWClient) Talent(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/talent/%d", options, id)
}

func (client *WoWClient) PvPTalentIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/pvp-talent/index", options)
}

func (client *WoWClient) PvPTalent(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/pvp-talent/%d", options, id)
}

// Tech Talent

func (client *WoWClient) TechTalentTreeIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/tech-talent-tree/index", options)
}

func (client *WoWClient) TechTalentTree(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/tech-talent-tree/%d", options, id)
}

func (client *WoWClient) TechTalentIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/tech-talent/index", options)
}

func (client *WoWClient) TechTalent(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/tech-talent/%d", options, id)
}

func (client *WoWClient) TechTalentMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/tech-talent/%d", options, id)
}

// Title

func (client *WoWClient) TitleIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/title/index", options)
}

func (client *WoWClient) Title(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/title/%d", options, id)
}

// WoW Token

func (client *WoWClient) WoWTokenIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, DYNAMIC_NS,  "/token/index", options)
}

// Character Profile

func (client *WoWClient) characterRequest(realmSlug string, characterSlug string, variation string, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(PROFILE, PROFILE_NS, fmt.Sprintf("/character/%s/%s%s", realmSlug, characterSlug, variation), options)
}

func (client *WoWClient) CharacterAchievements(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/achievements", options)
}

func (client *WoWClient) CharacterAchievementsStatistics(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/achievements/statistics", options)
}

func (client *WoWClient) CharacterAppearance(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/appearance", options)
}

func (client *WoWClient) CharacterCollections(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/collections", options)
}

func (client *WoWClient) CharacterMounts(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/collections/mounts", options)
}

func (client *WoWClient) CharacterPets(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/collections/pets", options)
}

func (client *WoWClient) CharacterEncounters(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/encounters", options)
}

func (client *WoWClient) CharacterDungeons(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/encounters/dungeons", options)
}

func (client *WoWClient) CharacterRaids(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/encounters/raids", options)
}

func (client *WoWClient) CharacterEquipment(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/equipment", options)
}

func (client *WoWClient) CharacterHunterPets(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/hunter-pets", options)
}

func (client *WoWClient) CharacterMedia(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/character-media", options)
}

func (client *WoWClient) CharacterMythicKeystoneProfile(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/mythic-keystone-profile", options)
}

func (client *WoWClient) CharacterMythicKeystoneSeasonDetails(realmSlug string, characterSlug string, seasonId int, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, fmt.Sprintf("/mythic-keystone-profile/season/%d", seasonId), options)
}

func (client *WoWClient) CharacterProfession(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/professions", options)
}

func (client *WoWClient) CharacterProfile(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "", options)
}

func (client *WoWClient) CharacterStatus(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/status", options)
}

func (client *WoWClient) CharacterPvP(realmSlug string, characterSlug string, bracket string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, fmt.Sprintf("/pvp-bracket/%s", bracket), options)
}

func (client *WoWClient) CharacterPvPSummary(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/pvp-summary", options)
}

func (client *WoWClient) CharacterQuests(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/quests", options)
}

func (client *WoWClient) CharacterCompletedQuests(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/quests/completed", options)
}

func (client *WoWClient) CharacterReputations(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/reputations", options)
}

func (client *WoWClient) CharacterSpecializations(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/specializations", options)
}

func (client *WoWClient) CharacterCompletedStatistics(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/statistics", options)
}

func (client *WoWClient) CharacterTitles(realmSlug string, characterSlug string, options *RequestOptions) *ApiResponse {
	return client.characterRequest(realmSlug, characterSlug, "/titles", options)
}

// Guild Profile

func (client *WoWClient) guildRequest(realmSlug string, guildSlug string, variation string, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, PROFILE_NS, fmt.Sprintf("/guild/%s/%s%s", realmSlug, guildSlug, variation), options)
}

func (client *WoWClient) GuildProfile(realmSlug string, guildSlug string, options *RequestOptions) *ApiResponse {
	return client.guildRequest(realmSlug, guildSlug, "", options)
}

func (client *WoWClient) GuildActivity(realmSlug string, guildSlug string, options *RequestOptions) *ApiResponse {
	return client.guildRequest(realmSlug, guildSlug, "/activity", options)
}

func (client *WoWClient) GuildAchievements(realmSlug string, guildSlug string, options *RequestOptions) *ApiResponse {
	return client.guildRequest(realmSlug, guildSlug, "/achievements", options)
}

func (client *WoWClient) GuildRoster(realmSlug string, guildSlug string, options *RequestOptions) *ApiResponse {
	return client.guildRequest(realmSlug, guildSlug, "/roster", options)
}
