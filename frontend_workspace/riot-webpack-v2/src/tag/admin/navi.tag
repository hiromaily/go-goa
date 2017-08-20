<navi>
  <div class="ui fixed inverted menu">
    <div class="ui container">
      <a each={ links } href="#{ url }" class={ header: name === "Admin", selected: parent.selectedId === url, item: true }>
        <img if={ name==='Admin' } class="logo" src="/img/hy.png" />
        { name }
      </a>
    </div>
  </div>

  <script>
    var self = this

    self.links = [
      { name: "Admin", url: "" },
      { name: "User", url: "user" },
      { name: "Company", url: "company" },
      { name: "Tech", url: "tech" }
    ]

    let r = riot.route.create()
    r(self.highlightCurrent)

    self.highlightCurrent = (id) => {
      self.selectedId = id
      self.update()
    }
  </script>
</navi>