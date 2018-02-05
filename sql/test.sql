INSERT INTO administrators(name, password, level) VALUES(
	'admin',
    'admin',
    0
);

INSERT INTO administrators(name, password, level) VALUES(
	'admin1',
    'admin1',
    0
);

INSERT INTO organizations(name, phone, password, collage, description) VALUES(
	'自嗨社',
    '15814092425',
    '5f4dcc3b5aa765d61d8327deb882cf99',
    '深圳市深圳中学',
    '简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介'
);

INSERT INTO organizations(name, phone, password, collage, description) VALUES(
	'麻将社',
    '110',
    '5f4dcc3b5aa765d61d8327deb882cf99',
    '广州中山大学',
    '这是打麻将的地方'
);

INSERT INTO organization_departments(name, organization_id) VALUES(
	'自嗨部门1',
    1
);

INSERT INTO organization_departments(name, organization_id) VALUES(
	'自嗨部门2',
    1
);

INSERT INTO activities(name, description, category, poster_url, logo_url,
	organization_id) VALUES(
	'长白山山顶自嗨活动',
    '丰富学生课余生活',
    '运动体育',
    null,
    null,
    1
);

INSERT INTO activities(name, description, category, poster_url, logo_url,
	organization_id) VALUES(
	'麻将大赛',
    '丰富学生课余生活',
    '益智',
    null,
    null,
    2
);

INSERT INTO activity_stages(stage_num, start_time, end_time, location, content,
	activity_id) VALUES(
	1,
    '2018-12-25 13:00:00',
    '2018-12-25 18:00:00',
    '广东省白山市',
    '点炮',
    2
);

INSERT INTO activity_stages(stage_num, start_time, end_time, location, content,
	activity_id) VALUES(
	1,
    '2018-12-25 13:00:00',
    '2018-12-25 18:00:00',
    '广东省白山市',
    '上长白山山顶唱我的中国心',
    1
);

INSERT INTO activity_stages(stage_num, start_time, end_time, location, content,
	activity_id) VALUES(
	2,
    '2018-12-25 13:00:00',
    '2018-12-25 18:00:00',
    '广东省白山市',
    '下长白山山脚唱最炫民族风',
    1
);

INSERT INTO users(username, phone, password, location, vip, avatar_url, camera,
	description, occupation, collage) VALUES(
	'test_user1',
    '15814092425',
    '5f4dcc3b5aa765d61d8327deb882cf99',
	'广东省广州市',
    0,
    null,
    '相机X',
    '个性签名个性签名个性签名个性签名个性签名个性签名个性签名',
    '在校学生',
    '中山大学'
);

INSERT INTO users(username, phone, password, location, vip, avatar_url, camera,
	description, occupation, collage) VALUES(
	'test_user1',
    '15814092425',
    '5f4dcc3b5aa765d61d8327deb882cf99',
	'广东省广州市',
    0,
    null,
    '相机Y',
    '签名个性签名个性签名个性签名个性签名个性签名个性签名个性签名个性',
    '在校学生',
    '深圳中学'
);

INSERT INTO photos(url, took_time, took_location, release_time,
	category, likes, reject_reason, photographer_id) VALUES(
    '/static/assets/tobee.png',
	'2019-1-25 13:00:00',
    '广州大山中学',
    '2019-1-25 14:00:00',
    '新闻',
    1,
    '无',
    1
);

INSERT INTO photos(url, took_time, took_location, release_time,
	category, likes, reject_reason, photographer_id) VALUES(
    '/static/assets/tobee.png',
	'2019-1-24 13:00:00',
    '广州华师',
    '2019-1-24 14:00:00',
    '微距',
    1,
    '无',
    1
);

INSERT INTO photos(url, took_time, took_location,
	category, likes, photographer_id) VALUES(
    '/static/assets/tobee.png',
	'2019-1-24 13:00:00',
    '未审核1',
    '未审核',
    1,
    1
);

INSERT INTO photos(url, took_time, took_location,
	category, likes, photographer_id) VALUES(
    '/static/assets/tobee.png',
	'2019-1-24 13:00:00',
    '未审核2',
    '未审核',
    1,
    1
);

INSERT INTO photos(url, took_time, took_location,
	category, likes, photographer_id) VALUES(
    '/static/assets/tobee.png',
	'2019-1-24 13:00:00',
    '未审核3',
    '未审核',
    1,
    1
);

INSERT INTO photos(url, took_time, took_location,
	category, likes, photographer_id) VALUES(
    '/static/assets/tobee.png',
	'2019-1-24 13:00:00',
    '未审核4',
    '未审核',
    1,
    1
);

INSERT INTO photo_tags(tag, photo_id) VALUES(	
	'风景',
    1
);

INSERT INTO photo_tags(tag, photo_id) VALUES(	
	'建筑',
    2
);

INSERT INTO photo_comments(title, content, user_id, photo_id) VALUES(
	'非常漂亮',
    '这照片美极了',
    1,
    1
);

INSERT INTO photo_comments(title, content, user_id, photo_id) VALUES(
	'难看',
    '这照片不好了',
    2,
    1
);

INSERT INTO photo_lives(expect_members, ad_progress, activity_stage_id,
	manager_id, photographer_manager_id) VALUES(
	1,
    '四饭-三饭',
	1,
    1,
    1
);

INSERT INTO photo_lives(expect_members, ad_progress, activity_stage_id,
	manager_id, photographer_manager_id) VALUES(
	2,
    '四饭-三饭',
	2,
    1,
    1
);

INSERT INTO user_login_logs(login_time, login_location, login_device,
	user_id) VALUES(
	'2019-1-24 13:00:00',
    '广州华师',
    'IPHONE8',
    1
);

INSERT INTO user_login_logs(login_time, login_location, login_device,
	user_id) VALUES(
	'2019-2-24 13:00:00',
    '广州华师',
    'IPHONE9',
    1
);

INSERT INTO organization_login_logs(login_time, login_location,
	login_device, organization_id) VALUES(
    '2019-1-24 13:00:00',
    '广州中山大学',
    'IPHONE8',
    1
);

INSERT INTO organization_login_logs(login_time, login_location,
	login_device, organization_id) VALUES(
    '2019-1-25 13:00:00',
    '广州中山大学东校区',
    'IPHONE9',
    1
);

INSERT INTO administrator_login_logs(login_time, login_location,
	login_device, administrator_id) VALUES(
    '2019-1-24 13:00:00',
    '广州华农',
    'IPHONE8',
    1
);

INSERT INTO administrator_login_logs(login_time, login_location,
	login_device, administrator_id) VALUES(
    '2019-1-26 13:00:00',
    '广州华农',
    'IPHONE9',
    1
);

INSERT INTO user_organization_relationships(user_id, organization_id) VALUES(
	1,
    1
);

INSERT INTO user_organization_relationships(user_id, organization_id) VALUES(
	2,
    2
);

INSERT INTO user_user_relationships(user_id, liked_user_id) VALUES(
	1,
    2
);

INSERT INTO user_photo_relationships(user_id, liked_photo_id) VALUES(
	1,
    1
);

INSERT INTO user_photo_relationships(user_id, liked_photo_id) VALUES(
	2,
    2
);

INSERT INTO user_activity_relationships(user_id, activity_id) VALUES(
	1,
    1
);

INSERT INTO user_activity_relationships(user_id, activity_id) VALUES(
	2,
    2
);


INSERT INTO organization_contact_relationships(organization_id, contact_id) VALUES(
	1,
    1
);

INSERT INTO organization_contact_relationships(organization_id, contact_id) VALUES(
	2,
    2
);

INSERT INTO photo_live_supervisor_relationships(photo_live_id, supervisor_id) VALUES(
	1,
    1
);

INSERT INTO photo_live_supervisor_relationships(photo_live_id, supervisor_id) VALUES(
	2,
    2
);