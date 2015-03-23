{{template "head.tpl" .}}

<div class="container theme-showcase" role="main">
	<div class="row">
		<div class="span12">
			<form id="login" method="post" action="#">
				<fieldset>
				    <legend>Login</legend> 

					<p style="color:#ff00ff">{{.warning}}</p>
					
                		    	<label>username</label><br>
                		    	<input name="username" type="text"/><br><br>

                		    	<label>Password</label><br>
					<input name="password" type="password"/><br><br>

					<button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
			</form>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
