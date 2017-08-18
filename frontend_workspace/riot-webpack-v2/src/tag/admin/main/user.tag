<user>
<div class="ui container" style="margin-bottom: 50px;">
    <h3 class="ui header">User Model</h3>

    <table class="ui celled striped table">
        <thead>
        <tr>
            <th>ID</th>
            <th>Name</th>
            <th>E-mail address</th>
            <th>Delete</th>
        </tr>
        </thead>
        <tbody>
        <tr>
            <td>1</td>
            <td>Hiroki Yasui</td>
            <td>hiroki@goa.com</td>
            <td class="collapsing">
                <button class="ui button">Delete</button>
            </td>
        </tr>
        <tr>
            <td>2</td>
            <td>Hiroki Yasui</td>
            <td>hiroki@goa.com</td>
            <td class="collapsing">
                <button class="ui button">Delete</button>
            </td>
        </tr>
        </tbody>
        <tfoot class="full-width">
        <tr>
            <th></th>
            <th>
                <div class="ui pagination menu">
                    <a class="item">1</a>
                    <div class="disabled item">...</div>
                    <a class="item active">10</a>
                    <a class="item">11</a>
                    <a class="item">12</a>
                </div>
            </th>
            <th colspan="2">
                <a href="adduser.html">
                    <div class="ui right floated small primary labeled icon button">
                        <i class="user icon"></i> Add User
                    </div>
                </a>
                <!--
                <div class="ui small button">
                    Approve
                </div>
                <div class="ui small  disabled button">
                    Approve All
                </div>
                -->
            </th>
        </tr>
        </tfoot>
    </table>

</div>

<div class="hidden ui container" style="margin-bottom: 100px;">
    <div class="ui"> <!-- style="display: none;" -->
        <h3 class="ui header">Add User</h3>

        <form class="ui fluid form">
            <div class="ui two column middle aligned very relaxed stackable grid">
                <div class="column">
                    <div class="ui form">
                        <div class="field">
                            <label>Username</label>
                            <div class="ui left icon input">
                                <input type="text" placeholder="Username">
                                <i class="user icon"></i>
                                <!--
                                <div class="ui left pointing label">
                                    That name is taken!
                                </div>
                                -->
                            </div>
                        </div>

                        <div class="field">
                            <label>Email</label>
                            <div class="ui left icon input">
                                <input type="text" placeholder="Email">
                                <i class="mail icon"></i>
                            </div>
                        </div>

                        <div class="field">
                            <label>Password</label>
                            <div class="ui left icon input">
                                <input type="text" placeholder="Password">
                                <i class="lock icon"></i>
                            </div>
                        </div>
                        <div class="ui blue submit button">Save</div>
                    </div>
                </div>
            </div>
        </form>
    </div>
</div>
</user>