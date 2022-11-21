<template>
    <div class="searchResultContainer">
        <search-item v-for="result in results" :key="result" :result="result"></search-item>
    </div>
</template>

<script>
import SearchItem from './SearchItem.vue'

export default{
    props:['results'],
    components:{SearchItem},
}
</script>

<style scoped>
.searchResultContainer{
    margin-top:15px;
    display:flex;
    height: calc(100% - 60px) ;
    flex-direction: column;
    overflow-y: auto;
    overscroll-behavior: contain;
    gap:10px;
}
span.highlight{
  background-color : yellow;
}
</style>