<template>
    <div class="searchItem">
        <p class="searchResultAuthor">{{result.item.sender.nickname}}</p>
        <div v-html="higlighted"></div>
    </div>
</template>

<script>
export default{
    props:['result'],
    data(){
    return{
        highlightClass: "highlight"
    }
   },
    computed:{
        higlighted(){
            return this.higlight()
        }
    },
    methods:{
        higlight(){
            let content ="";
            let nextUnhighlightedRegionStartingIndex = 0;
            this.result.matches[0].indices.forEach(region => {
                const lastRegionNextIndex = region[1] + 1;

                content += [
                this.result.item.content.substring(nextUnhighlightedRegionStartingIndex, region[0]),
                `<span class="${this.highlightClass}" style="background-color:#4f34927e">`,
                this.result.item.content.substring(region[0], lastRegionNextIndex),
                '</span>',
            ].join('');

            nextUnhighlightedRegionStartingIndex = lastRegionNextIndex;
            })
             content += this.result.item.content.substring(nextUnhighlightedRegionStartingIndex);
                console.log("Res:", content)
             return content;
            }
    }
}
</script>

<style scoped>
.searchItem{
    padding: 8px;
    border-radius: 10px;
    word-break: break-word;
    display: inline-block;
    background-color: var(--input-bg);
    
}
.searchItem div{
font-size:14px;
    
}
.searchResultAuthor{
    color:var(--color-blue);
    font-size: 10px;
}
</style>