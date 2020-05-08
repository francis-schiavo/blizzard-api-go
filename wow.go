package blizzard_api

import (
	"github.com/go-redis/redis"
	"net/http"
)

type WoWClient struct {
	ApiClient
}

func NewWoWClient(region Region, redisClient *redis.Client, validCacheStatus []int) *WoWClient {
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

// Mount

func (client *WoWClient) MountIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/mount/index", options)
}

func (client *WoWClient) Mount(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/mount/%d", options, id)
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

// Title

func (client *WoWClient) TitleIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/title/index", options)
}

func (client *WoWClient) Title(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/title/%d", options, id)
}