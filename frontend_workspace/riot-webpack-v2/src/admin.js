import route from 'riot-route'
riot.route = route

require('./tag/admin/navi.tag');
require('./tag/admin/header.tag');
require('./tag/admin/main.tag');
require('./tag/admin/main/admin.tag');
require('./tag/admin/main/company.tag');
require('./tag/admin/main/tech.tag');
require('./tag/admin/main/user.tag');

// route('/', () => riot.mount('main', 'home'))
// route('/about', () => riot.mount('main', 'about'))
// route('/contact', () => riot.mount('main', 'contact'))

riot.mount('*')
riot.route.start(true)
