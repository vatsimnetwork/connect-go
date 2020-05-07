package Layout

// Generate function generates application layout
func Generate(css, html, script string) string {
	return "<html>" +
		"<head>" +
		"<!-- Font Awesome -->\n<link rel=\"stylesheet\" href=\"https://use.fontawesome.com/releases/v5.8.2/css/all.css\">" +
		"<!-- Google Fonts -->" +
		"<link rel=\"stylesheet\" href=\"https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap\">" +
		"<!-- Bootstrap core CSS -->\n<link href=\"https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.4.1/css/bootstrap.min.css\" rel=\"stylesheet\">\n<!-- Material Design Bootstrap -->\n<link href=\"https://cdnjs.cloudflare.com/ajax/libs/mdbootstrap/4.16.0/css/mdb.min.css\" rel=\"stylesheet\">"+
		css +
		"<body>" +
		html+
		"<!-- JQuery -->\n<script type=\"text/javascript\" src=\"https://cdnjs.cloudflare.com/ajax/libs/jquery/3.4.1/jquery.min.js\"></script>\n<!-- Bootstrap tooltips -->\n<script type=\"text/javascript\" src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.4/umd/popper.min.js\"></script>\n<!-- Bootstrap core JavaScript -->\n<script type=\"text/javascript\" src=\"https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.4.1/js/bootstrap.min.js\"></script>\n<!-- MDB core JavaScript -->\n<script type=\"text/javascript\" src=\"https://cdnjs.cloudflare.com/ajax/libs/mdbootstrap/4.16.0/js/mdb.min.js\"></script>" +
		script +
		"</body>" +
		"</html>"
}