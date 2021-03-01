import route from 'riot-route'
riot.route = route

require('./tag/admin/navi.tag');
require('./tag/admin/header.tag');
require('./tag/admin/main.tag');
require('./tag/admin/main/admin.tag');
require('./tag/admin/main/company.tag');
require('./tag/admin/main/tech.tag');
require('./tag/admin/main/user.tag');

require('./tag/admin/main/user_detail.tag');

require('./tag/resume/like_tech.tag');
require('./tag/resume/dislike_tech.tag');
require('./tag/resume/work_history.tag');
require('./tag/common/raw.tag');


riot.mount('*')
riot.route.start(true)
