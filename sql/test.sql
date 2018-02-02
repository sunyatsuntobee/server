INSERT INTO administrators(name, password, level) VALUES(
	'admin',
    'admin',
    0
);

INSERT INTO organizations(name, phone, password, collage, description) VALUES(
	'自嗨社',
    '15814092425',
    '5f4dcc3b5aa765d61d8327deb882cf99',
    '深圳市深圳中学',
    '简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介'
);

INSERT INTO organization_departments(name, organization_id) VALUES(
	'自嗨部门1',
    1
);

INSERT INTO organization_departments(name, organization_id) VALUES(
	'自嗨部门2',
    1
);

INSERT INTO activities(name, description, category, poster_url, logo,
	organization_id) VALUES(
	'长白山山顶自嗨活动',
    '丰富学生课余生活',
    '运动体育',
    null,
    null,
    1
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
