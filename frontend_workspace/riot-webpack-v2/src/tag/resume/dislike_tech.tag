<dislike-tech>
    <div class="ui segments">
        <div class="ui segment">
            Dislike <i if={this.edit} class="edit icon" style="padding-left:10px;"></i>
        </div>
        <div class="ui segments">
            <div each="{ opts.items }" class="ui segment">
                <p>{ tech_name }</p>
            </div>
        </div>
    </div>

<script>
self = this
self.edit = false
if (location.pathname.match(/admin/)){
  self.edit = true
}
</script>

</dislike-tech>
