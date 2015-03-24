<!DOCTYPE html>
<html>
    <head>
        <title>GDCPC Register System</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <link rel="stylesheet" href="/static/css/bootstrap.min.css">
        <link rel="stylesheet" href="/static/css/bootstrap-theme.min.css">
        <link rel="stylesheet" href="/static/css/theme.css">
		<script type="text/javascript" src="/static/jquery/jquery-2.1.3.min.js"></script>
    </head>

    <body role="document">
        <nav class="navbar navbar-inverse navbar-fixed-top">
            <div class="container">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
                        <span class="sr-only">Toggle navigation</span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                    <a class="navbar-brand" href="/index">GDCPC Register System</a>
                </div>
                <div class="navbar-collapse collapse">
                    <ul class="nav navbar-nav">
                        <li><a href="/index">Home</a></li>
			{{ if $.logined }}
				<li><a href="/show_file">Personal Profile</a></li>
                        	<li><a href="/create_team">Create Team</a></li>
                        	<li><a href="/edit_coach">Setting</a></li>
			{{ end }}
                    </ul>
                    <ul class="nav navbar-nav pull-right">
			{{ if $.logined }}
				<li><a href="/show_file">{{.username}}</a></li>
				<li><a href="/logout">Logout</a></li>
			{{ else }}
                  		<li><a href="/login">Login</a></li>
                        	<li><a href="/reg_coach">Register</a></li>
			{{ end }}
                    </ul>
                </div><!--/.nav-collapse -->
            </div>
        </nav>
