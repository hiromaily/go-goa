<user_detail>
<div class="ui container" style="margin-bottom: 50px;">


  <div class="ui two column grid">

    <!-- start left column -->
    <div class="four wide column">
      <like-tech></like-tech>

      <br/>

      <dislike-tech></dislike-tech>

    </div>
    <!-- end left column -->

    <!-- start right column-->
    <div class="twelve wide column">

      <work-history></work-history>

    </div>
    <!-- end right column -->

  </div>
</div>
<script>
//this.opts.items = this.parent.items
self = this

//set resume page
console.log('this.parent.user_id:', this.parent.user_id)
rg.setResume(riot, this.parent.user_id)

</script>
</user_detail>