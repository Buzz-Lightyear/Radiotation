<template>
  <div class="track">
    <div class="columns is-gapless is-mobile">
      <div class="thumbnail column is-2">
        <img :src="image">
      </div>
      <div class="column is-10">
        <div class="metadata-holder">
          <div class="title metadata">{{Name}}</div>
          <div class="artist metadata">{{artist}}</div>
          <div class="album metadata">{{Album.Name}}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Track',
  props: ['Artists', 'Name', 'ID', 'Album'],
  computed: {
    image () {
      var url = 'https://via.placeholder.com/150x150'
      if (this.Album.Images.length > 0) {
        url = this.Album.Images[0].URL
      }
      return url
    },
    artist () {
      var names = []
      for (const artist of this.Artists) {
        names.push(artist.Name)
      }
      return names.join(', ')
    }
  }
}
</script>

<style scoped>
.track {
  padding-top: 6px;
  padding-bottom: 6px;
}

.thumbnail {
  padding: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.thumbnail img {
  max-width: 100%;
}

.metadata-holder {
  padding-left: 12px;
}

.metadata {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  padding: 0;
  margin: 0;
}

.title {
  font-size: 16px;
  font-weight: bold;
}

.artist {
  font-size: 12px;
}

.album {
  font-size: 10px;
}
</style>
