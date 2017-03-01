{{define "title"}}Search{{end}}

{{define "content"}}
<div class="container">
    <div class="row align-items-center">
        <div class="col-lg-6 mx-lg-auto">
            <a href="/template/home.tpl">
                <img src="/static/img/logo/dole-logo.png" height="256" class="mx-auto d-block logo" alt="Dole Logo">
            </a>
            <form id="search-form">
                <input type="text" class="form-control form-control-lg" placeholder="Search applicants">
            </form>
        </div>
    </div>
</div>
{{end}}

