<like-tech>
    <div class="ui segments">
        <div class="ui segment">
            Like <i if={this.edit} class="edit icon" style="padding-left:10px;"></i>
        </div>
        <div class="ui segments">
            <div each="{ opts.items }" class="ui segment">
                <p>{ tech_name }</p>
            </div>
        </div>
    </div>

<script>
self = this
//After mount, this parent can not be access
//console.log(this.parent)
//console.log('this.parent.edit:', this.parent.edit)

//if url includes admin, set edit
self.edit = false
if (location.pathname.match(/admin/)){
  self.edit = true
}
</script>

</like-tech>

