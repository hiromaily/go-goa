<work-history>
    <div each="{ opts.items }" class="ui card">
        <div class="content" style="background: #787878;">
            <div class="header" style="color: #FFFFFF;">
                { title } at { company }
                <i class="{ coutry} flag" style="padding-left:10px;"></i>
            </div>
        </div>
        <div class="content">
            <h4 class="ui sub header">{ term }</h4>
            <div class="ui small feed">
                <div class="event">
                    <div class="content">
                        <div class="summary">
                            <ul class="ui list">
                                <li each="{ key in description }">
                                  { key }
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="extra content" style="background:rgb(243, 244, 245);">
            <div each="{ techs }" class="ui black basic button">
              { body }
            </div>
        </div>
    </div>

    <div class="ui section divider"></div>

</work-history>

