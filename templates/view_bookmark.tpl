{{ define "content" }}
<div class="content">
    <h4 class="title">
        <span class="icon-text">
            {{ if .Bookmark.Favicon }}
            <span class="icon">
                <img src="{{ .Bookmark.Favicon | ToURL }}" alt="favicon" />
            </span>
            {{ end }}
            <span>
                <a href="{{ .Bookmark.URL }}" target="_blank">{{ .Bookmark.Title }}</a>
            </span>
        </span>
        <p class="is-size-7 has-text-grey has-text-weight-normal">{{ Truncate .Bookmark.URL 100 }}</p>
    </h4>
    <p>{{ .Bookmark.Notes }}</p>
    {{ $uid := 0 }}
    {{ if .User }}{{ $uid = .User.ID }}{{ end }}
    {{ if .Bookmark.Tags }}
        {{ range .Bookmark.Tags }}
        <a href="{{ if ne $uid $.Bookmark.UserID }}{{ BaseURL "/bookmarks" }}{{ else }}{{ BaseURL "/my_bookmarks" }}{{ end }}?tag={{ .Text }}"><span class="tag is-info">{{ .Text }}</span></a>
        {{ end }}
    {{ end }}
    {{ block "snapshots" KVData "Snapshots" .Bookmark.Snapshots "IsOwn" (eq .Bookmark.UserID $uid ) }}{{ end }}
    {{ .Bookmark.CreatedAt | ToDate }} {{ if .Bookmark.Public }}Public{{ else }}Private{{ end }}
    {{ if .User }}
      {{ if eq .User.ID .Bookmark.UserID }}
        <a href="{{ BaseURL "/edit-bookmark" }}?id={{ .Bookmark.ID }}">edit</a>
      {{ end }}
    {{ end }}
</div>
{{ end }}
