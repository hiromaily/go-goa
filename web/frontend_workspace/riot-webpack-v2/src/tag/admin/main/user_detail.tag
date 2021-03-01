<user_detail>
<div class="ui container" style="margin-bottom: 50px;">

  <h3 class="ui header">
    <img src="/img/hy.png" class="ui circular image">
    <a href="https://hiromaily.github.io/" style="color:#333333;">{this.parent.items.user_name}</a>
  </h3>
  <div style="padding-bottom: 30px; padding-left: 30px">
    <p>I'm working as software engineer for 14 years and came from Japan to join multinational environment in Oct 2016.</p>
    <p>And I am looking for job which requires Golang. However I always have to require working VISA because of my nationality. Thank you.</p>
  </div>

  <div class="ui two column grid">

    <!-- start left column -->
    <div class="four wide column">
      <like-tech if={ this.parent.tag==='user_detail' } />

      <br/>

      <dislike-tech if={ this.parent.tag==='user_detail' } />

    </div>
    <!-- end left column -->

    <!-- start right column-->
    <div class="twelve wide column">

      <work-history if={ this.parent.tag==='user_detail' } />

    </div>
    <!-- end right column -->

  </div>
</div>
<script>
//this.opts.items = this.parent.items
self = this

//set resume page
//console.log('this.parent.user_id:', this.parent.user_id)
rg.setResume(riot, this.parent.user_id)

</script>
</user_detail>