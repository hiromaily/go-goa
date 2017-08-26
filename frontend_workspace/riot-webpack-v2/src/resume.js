require('./tag/resume/like_tech.tag');
require('./tag/resume/dislike_tech.tag');
require('./tag/resume/work_history.tag');
require('./tag/common/raw.tag');


function setResume(){
    // data definition
    let data = [
        {url: `/api/user/${sessionStorage.getItem('id')}/liketech`, element: 'like-tech'},
        {url: `/api/user/${sessionStorage.getItem('id')}/disliketech`, element: 'dislike-tech'},
        {url: `/api/user/${sessionStorage.getItem('id')}/workhistory`, element: 'work-history'}
    ];
    if (window.debugMode == 1){
        data[0].url='/json/liketech.json';
        data[1].url='/json/disliketech.json';
        data[2].url='/json/workhistory.json';
    }


    //loop
    data.forEach((obj) => {
        let fn = (json, element) => {
            if(json.status && json.status != 200){
                sessionStorage.removeItem('jwt');
                sessionStorage.removeItem('id');
                location.href = '/login.html';
                return;
            }
            //Success
            riot.mount(element, {
                items : json
            } );
        }

        rg.callAPI(obj.url, {}, 'GET', fn, obj.element)
    });

}

setResume();
