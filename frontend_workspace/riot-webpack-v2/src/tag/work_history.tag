<work-history>
<virtual each="{ opts.items }">
    <div  class="ui card">
        <div class="content" style="background: #787878;">
            <div class="header" style="color: #FFFFFF;">
                <virtual if={company=='Freelancer'}>{ title } as { company }</virtual>
                <virtual if={company!='Freelancer'}>{ title } at { company }</virtual>
                <i class="{ country} flag" style="padding-left:10px;"></i>
            </div>
        </div>
        <div class="content">
            <h4 class="ui sub header">{ term }</h4>
            <div class="ui small feed">
                <div class="event">
                    <div class="content">
                        <div class="summary">
                            <ul class="ui list">
                                <virtual each="{ desc in description }">
                                  <li if={ isString(desc) }><raw content={ desc } /></li>
                                  <ul if={ isArray(desc) }>
                                    <li each="{ v in desc }">
                                      <raw content={ v } />
                                    </li>
                                  </ul>
                                </virtual>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="extra content" style="background:rgb(243, 244, 245);">
            <div each="{ tech in techs }" class="ui black basic button">
              { tech }
            </div>
        </div>
    </div>

    <div class="ui section divider"></div>
</virtual>


<script>
    isArray(obj) {
        return riot.util.check.isArray(obj)
    }
    isString(obj) {
        return riot.util.check.isString(obj)
    }

    //this.on('mount', () => {
        //console.log('finished');
        //remove parents
        //$(this).unwrap();
    //})

</script>

</work-history>

