namespace go api

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
    6: string avatar
    7: string background_image
    8: string signature
    9: i64 total_favorited
    10: i64 work_count
    11: i64 favorite_count
}

struct Video {
    1: i64 id              // 视频唯一标识
    2: User author           // 视频作者信息
    3: string play_url       // 视频播放地址
    4: string cover_url      // 视频封面地址
    5: i64 favorite_count  // 视频的点赞总数
    6: i64 comment_count   // 视频的评论总数
    7: bool is_favorite      // true-已点赞，false-未点赞
    8: string title          // 视频标题
}

struct Comment {
    1: i64 id            // 视频评论id
    2: User user           // 评论用户信息
    3: string content      // 评论内容
    4: string create_date  // 评论发布日期，格式 mm-dd
}

struct douyin_feed_request {
    1: i64 latest_time  // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: string token       // 可选参数，登录用户设置
}

struct douyin_feed_response {
    1: i32 status_code       // 状态码，0-成功，其他值-失败
    2: string status_msg       // 返回状态描述
    3: list<Video> video_list  // 视频列表
    4: i64 next_time         // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct douyin_user_register_request {
    1: string username (api.query="username", api.vd="len($) > 0")
    2: string password (api.query="password", api.vd="len($) > 0")
}

struct douyin_user_register_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败
    2: string status_msg  // 返回状态描述
    3: i64 user_id      // 用户id
    4: string token       // 用户鉴权token
}

struct douyin_user_login_request {
    1: string username (api.query="username", api.vd="len($) > 0")
    2: string password (api.query="password", api.vd="len($) > 0")
}


struct douyin_user_login_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败
    2: string status_msg  // 返回状态描述
    3: i64 user_id      // 用户id
    4: string token       // 用户鉴权token
}

struct douyin_user_request {
    1: i64 user_id (api.query="user_id")
    2: string token (api.query="token", api.vd="len($) > 0")
}

struct douyin_user_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败
    2: string status_msg  // 返回状态描述
    3: User user          // 用户信息
}

struct douyin_publish_action_request {
    1: string token  (api.form="token", api.vd="len($) > 0")// 用户鉴权token
    2: binary data    (api.form="data")// 视频数据
    3: string title  (api.form="title", api.vd="len($) > 0")// 视频标题
}

struct douyin_publish_action_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败
    2: string status_msg  // 返回状态描述
}

struct douyin_publish_list_request {
    1: i64 user_id  (api.query="user_id")// 用户id
    2: string token (api.query="token", api.vd="len($) > 0")// 用户鉴权token
}

struct douyin_publish_list_response {
    1: i32 status_code       // 状态码，0-成功，其他值-失败
    2: string status_msg       // 返回状态描述
    3: list<Video> video_list  // 用户发布的视频列表
}

struct douyin_favorite_action_request {
    1: string token     (api.query="token", api.vd="len($) > 0")  // 用户鉴权token
    2: i64 video_id     (api.query="video_id")// 视频id
    3: i32 action_type  (api.query="action_type")// 1-点赞，2-取消点赞
}

struct douyin_favorite_action_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败
    2: string status_msg  // 返回状态描述
}

struct douyin_favorite_list_request {
    1: i64 user_id    (api.query="user_id")// 用户id
    2: string token   (api.query="token", api.vd="len($) > 0")// 用户鉴权token
}

struct douyin_favorite_list_response {
    1: i32 status_code       // 状态码，0-成功，其他值-失败
    2: string status_msg       // 返回状态描述
    3: list<Video> video_list  // 用户点赞视频列表
}

struct douyin_comment_action_request {
    1: string token        (api.query="token", api.vd="len($) > 0") // 用户鉴权token
    2: i64 video_id        (api.query="video_id")// 视频id
    3: i32 action_type     (api.query="action_type")// 1-发布评论，2-删除评论
    4: string comment_text (api.query="comment_text", api.vd="len($) > 0")// 用户填写的评论内容，在action_type=1的时候使用
    5: i64 comment_id      (api.query="comment_id")// 要删除的评论id，在action_type=2的时候使用
}

struct douyin_comment_action_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败
    2: string status_msg  // 返回状态描述
    3: Comment comment    // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct douyin_comment_list_request {
    1: string token    (api.query="token", api.vd="len($) > 0")// 用户鉴权token
    2: i64 video_id    (api.query="video_id")// 视频id
}

struct douyin_comment_list_response {
    1: i32 status_code           // 状态码，0-成功，其他值-失败
    2: string status_msg           // 返回状态描述
    3: list<Comment> comment_list  // 评论列表
}

struct douyin_relation_action_request {
    1: string token       (api.query="token", api.vd="len($) > 0")// 用户鉴权token
    2: i64 to_user_id     (api.query="to_user_id")// 对方用户id
    3: i32 action_type    (api.query="action_type")// 1-关注，2-取消关注
}

struct douyin_relation_action_response {
    1: i32 status_code  // 状态码，0-成功，其他值-失败
    2: string status_msg  // 返回状态描述
}

struct douyin_relation_follow_list_request {
    1: i64 user_id    (api.query="user_id")// 用户id
    2: string token   (api.query="token", api.vd="len($) > 0")// 用户鉴权token
}

struct douyin_relation_follow_list_response {
    1: i32 status_code     // 状态码，0-成功，其他值-失败
    2: string status_msg     // 返回状态描述
    3: list<User> user_list  // 用户信息列表
}

struct douyin_relation_follower_list_request {
    1: i64 user_id   (api.query="user_id")// 用户id
    2: string token  (api.query="token", api.vd="len($) > 0")// 用户鉴权token
}

struct douyin_relation_follower_list_response {
    1: i32 status_code     // 状态码，0-成功，其他值-失败
    2: string status_msg     // 返回状态描述
    3: list<User> user_list  // 用户列表
}

service ApiService {
    douyin_feed_response Feed(1: douyin_feed_request req) ( api.post = "/douyin/feed/" )
    douyin_user_register_response DouyinUserRegister(1: douyin_user_register_request req) (api.post="/douyin/user/register/")
    douyin_user_login_response DouyinUserLogin(1: douyin_user_login_request req) (api.post="/douyin/user/login/")
    douyin_user_response DouyinUserGet(1: douyin_user_request req) (api.get="/douyin/user/")
    douyin_publish_action_response Publish(1: douyin_publish_action_request req) ( api.post = "/douyin/publish/action/" )
    douyin_publish_list_response GetPublishList(1: douyin_publish_list_request req) ( api.get = "/douyin/publish/list/" )
    douyin_favorite_action_response Favorite(1: douyin_favorite_action_request req) ( api.post = "/douyin/favorite/action/" )
    douyin_favorite_list_response GetFavoriteList(1: douyin_favorite_list_request req) ( api.get = "/douyin/favorite/list/" )
    douyin_comment_action_response Comment(1: douyin_comment_action_request req) ( api.post = "/douyin/comment/action/" )
    douyin_comment_list_response GetCommentList(1: douyin_comment_list_request req) ( api.get = "/douyin/comment/list/ " )
    douyin_relation_action_response Subscribe(1: douyin_relation_action_request req) ( api.post = "/douyin/relation/action/" )
    douyin_relation_follow_list_response GetFollowList(1: douyin_relation_follow_list_request req) ( api.get = "/douyin/relation/follow/list/" )
    douyin_relation_follower_list_response GetFollowerList(1: douyin_relation_follower_list_request req) ( api.get = "/douyin/relation/follower/list/" )
}