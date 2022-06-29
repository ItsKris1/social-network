<template>
    <NavBarOn />
    <NewPost />
    <AllPosts />
<<<<<<< HEAD


    <div class="item-list">
        <h3>Title</h3>
        <ul class="item-list__items">
            <li class="item-list__item">
                <div class="user-picture small"></div>
                <div class="item-list__text">Item 1</div>
            </li>
            <li class="item-list__item">
                <div class="user-picture small"></div>
                <div class="item-list__text">Item 2</div>
            </li>

            <li class="item-list__item">
                <div class="user-picture small"></div>
                <div class="item-list__text">Item 3</div>
            </li>

            <li class="item-list__item">
                <div class="user-picture small"></div>
                <div class="item-list__text">Item 4</div>
            </li>
        </ul>
        <button class="btn">Button<i class="uil uil-arrow-circle-right "></i></button>
    </div>

    <div class="item-list">
        <h3>Title</h3>
        <ul class="item-list__items">
            <li class="item-list__item">
                <img class="small" src="../assets/arrow-circle-right.svg">
                <div class="item-list__text">Item 1</div>
            </li>
            <li class="item-list__item">
                <img class="small" src="../assets/arrow-circle-right.svg">
                <div class="item-list__text">Item 2</div>
            </li>
            <li class="item-list__item">
                <img class="small" src="../assets/arrow-circle-right.svg">
                <div class="item-list__text">Item 3</div>
            </li>
            <li class="item-list__item">
                <img class="small" src="../assets/arrow-circle-right.svg">
                <div class="item-list__text">Item 4</div>
            </li>

        </ul>
        <button class="btn">Button<i class="uil uil-arrow-circle-right"></i></button>
    </div>
=======
    <Groups />
>>>>>>> origin/vue
</template>

<script>
import NavBarOn from '@/components/NavBarOn.vue'
import NewPost from '@/components/NewPost.vue'
import AllPosts from '@/components/AllPosts.vue'
import Groups from '@/components/Groups.vue'

export default {
    name: 'MainView',
<<<<<<< HEAD
    components: { NavBarOn, AllPosts },
=======
    components: { NavBarOn, NewPost, AllPosts, Groups },
>>>>>>> origin/vue
}
</script>

<style>
.item-list {
    display: flex;
    flex-direction: column;
    gap: 25px;
    padding: 25px 30px 25px 30px;
    background-color: var(--color-white);
    width: max-content;
    max-width: 250px;
    box-shadow: var(--container-shadow);
    border-radius: 10px;

}


.item-list__item {
    display: flex;
    gap: 7.5px;
    align-items: center;

}

.item-list__item:nth-child(1)~* {
    margin-top: 15px;
}

.item-list__text {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}
</style>