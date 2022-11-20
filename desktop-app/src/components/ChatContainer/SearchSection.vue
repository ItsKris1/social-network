<template>
    <div class="search-container" :class="toggleClass">
        <div class="search-btn" @click="toggleSearch" :class="toggleClass"></div>
        <input v-if="sectionOpen" type="text" class="search-input" :class="toggleClass"/>
        <!-- <div v-if="sectionOpen" class="searchResultContainer">Search results</div> -->
    </div>
</template>

<script>
export default{
   data(){
    return{
        initialOpen:true,
        sectionOpen: false
    }
   },
   computed:{
    toggleClass(){
        if (!this.sectionOpen && this.initialOpen){
            this.toggleInitialClass()
            return ""
        }else if (this.sectionOpen){
            return "open"
        }else{
            return "close"
        }
    }
   },
   methods:{
    toggleSearch(){
        this.sectionOpen = !this.sectionOpen;
    },
    toggleInitialClass(){
         this.initialOpen = false
    }
   }
}
</script>

<style scoped>
.search-container{
    position: relative;
    padding: 0;
    width:0;
}
/* .search-container.close{
    animation: 2s 1 forwards slideout;
}
.search-container.open{
    animation: 2s 1 forwards slidein;
} */
@media only screen and (min-width: 1100px) {
    /* styles for browsers larger than 960px; */
    .search-container.close{
        animation: 2s 1 forwards slideout;
    }
    .search-container.open{
        animation: 2s 1 forwards slidein;
    }
}
@media only screen and (max-width: 1100px) {
    /* styles for browsers smaller than 1100px; */
    .search-container.close{
        padding: 0;
        width:0;
    }
    .search-container.open{
        padding: 15px;
        width:30%;
    }
}
@keyframes slideout {
    from{
        padding: 15px;
        width:30%;
    }
    to{
        padding: 0;
        width:0;

    }
}
@keyframes slidein {
    from{
        padding: 0;
        width:0;
    }
    to{
        padding: 15px;
        width:500px;
    }
}
.search-btn{
    position:absolute;
    top:25px;
    background-image: url('../../assets/icons/glass.png');
    width: 20px;
    height:20px;
    cursor:pointer;
    left:-45px;
}
/* .search-btn.close{
    animation: 2s 1 forwards slideout-btn;
} */
@media only screen and (min-width: 1100px) {
    /* styles for browsers larger than 960px; */
    .search-btn.close{
        animation: 2s 1 forwards slideout-btn;
    }
    .search-btn.open{
        animation: 2s 1 forwards slidein-btn;
    }

}
@media only screen and (max-width: 1100px) {
    /* styles for browsers smaller than 1100px; */
.search-btn.close{
        left:-45px;
    }
    .search-btn.open{
        left:15px;
        display:inline-block;
    }
}
@keyframes slideout-btn {
    from{
        left:15px;
        display:inline-block;
    }
    to{
        left:-45px;
    }
}
/* .search-btn.open{
    animation: 2s 1 forwards slidein-btn;
} */
@keyframes slidein-btn {
    from{
        left:-45px;
    }
    to{
        left:15px;
        display:inline-block;
    }
}
/* .searchResultContainer{
    width: 300px;
} */

.search-input{
    border-radius:10px;
    margin-left:35px;
    box-shadow: var(--container-shadow);
    background-color:var(--input-bg);
}
/* .search-input.open{
    animation: 2s 1 forwards slidein-input;
} */
@media only screen and (min-width: 1100px) {
    /* styles for browsers larger than 960px; */
    .search-input.open{
    animation: 2s 1 forwards slidein-input;
}

}
@media only screen and (max-width: 1100px) {
    /* styles for browsers smaller than 1100px; */
    .search-input.open{
     width:80%;
    padding:10px;
    display:inline-block;
}
}
@keyframes slidein-input {
    from{
        /* margin-left:35px; */
        width:0px;
        padding:0px;
        display:none;
    }
    to{
        /* margin-left:35px; */
        width:80%;
        padding:10px;
        display:inline-block;
    }
}
</style>
