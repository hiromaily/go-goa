require('./tag/like_tech.tag');
require('./tag/dislike_tech.tag');
require('./tag/work_history.tag');


//Ajax
function callAPI(obj) {
    //like-tech
    fetch(obj.url, {
        headers: { 'Authorization': 'Bearer: xxxxx' }
    }).then(function(response) {
        return response.json();
    }).then(function(json) {
        riot.mount(obj.element, {
            items : json.items
        } );
    });
}

// data definition
let data = [
    {url: '/json/liketech.json', element: 'like-tech'},
    {url: '/json/disliketech.json', element: 'dislike-tech'},
    {url: '/json/workhistory.json', element: 'work-history'}
];

//loop
data.forEach(function(obj){
    callAPI(obj);
});


/*
//dislike-tech
fetch('/json/disliketech.json', {
    headers: { 'Authorization': 'Bearer: xxxxx' }
}).then(function(response) {
    return response.json();
}).then(function(json) {
    // riot.mount('dislike-tech', {
    //     items : [
    //         { tech : 'Microsoft .NET'},
    //         { tech : 'Ruby'},
    //         { tech : 'Java'}
    //     ]
    // } );
    riot.mount('dislike-tech', {
        items : json
    } );
});

//work-history
riot.mount('work-history', {
    items : [
        {
            title: 'Full Stack Engineer',
            company: '2Gears',
            country: 'nl',
            term: 'Dec 2016 - Jul 2017',
            description: ['Developed fin-tech web application both front-end and back-end.', 'Setting <a>Docker</a> configuration.'],
            techs: ['Node.js', 'ES6', 'Ext JS', 'Docker', 'PostgreSQL', 'JIRA']
        },
        {
            title: 'Full Stack Engineer',
            company: '2Gears',
            country: 'nl',
            term: 'Dec 2016 - Jul 2017',
            description: ['Developed fin-tech web application both front-end and back-end.', 'Setting <a>Docker</a> configuration.'],
            techs: ['Node.js', 'ES6', 'Ext JS', 'Docker', 'PostgreSQL', 'JIRA']
        }
    ]
} );
*/