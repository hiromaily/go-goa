require('./tag/like_tech.tag');
require('./tag/dislike_tech.tag');
require('./tag/work_history.tag');
require('./tag/raw.tag');


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
    {url: '/json/workhistory.json', element: 'work-history'},
];


//loop
data.forEach(function(obj){
    callAPI(obj);
});
