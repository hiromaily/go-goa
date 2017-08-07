<app>
   <h2>{ opts.title }</h2>
   <button onclick="{ show_count }">click</button>

   <script>
      self = this
      let count = 0

      show_count() {
         count += 1
         alert(`count: ${count}`)
      }
   </script>
</app>
