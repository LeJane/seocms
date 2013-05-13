<input type="hidden" name="location" value="{{.Category.NameEn}}" />
<h1>{{.Category.Name}}相关文章</h1>
<table class="article-list table table-striped table-hover">
<thead>
    <tr><th>文章标题</th><th>发表日期</th></tr>
</thead>
<tbody>
    {{$category := .Category.NameEn}}
    {{range .Articles}}
    <tr><td><a href="/{{$category}}/{{.Id}}" target="_blank">{{.Title}}</a></td><td>{{.Pubdate}}</td></tr>
    {{end}}
</tbody>
</table>