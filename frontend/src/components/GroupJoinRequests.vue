<template>
    <div v-if="isAdmin && this.joinRequests.length > 0" class="right-section">
        <div class="item-list__wrapper" id="requests">
            <h3>Requests</h3>
            <ul class="item-list">
                <div v-for="user in this.joinRequests">
                    <li>
                        <div class="row1">
                            <img class="small" :src="'http://localhost:8081/' + user.avatar">
                            <div>
                                <span class="who">{{ user.nickname }}</span>
                            </div>
                        </div>
                        <div class="row2">
                            <i class="uil uil-times accept" @click="reactToRequest(user.id,'accept')"></i>
                            <i class="uil uil-check decline" @click="reactToRequest(user.id,'decline')"></i>
                        </div>
                    </li>
                </div>
            </ul>
        </div>
    </div>
</template>

<script>

export default {
    name: "GroupJoinRequests",
    props: {
        isAdmin: false
    },
    data() {
        return {
            joinRequests: []
        }
    },
    created() {
        if (this.isAdmin) {
            this.getJoinRequests()
        }
    },
    methods: {
        async getJoinRequests() {
            await fetch("http://localhost:8081/groupRequests?groupId=" + this.$route.params.id, {
                credentials: "include"
            }).then(r => r.json()).then(json => {
                console.log(json)
                this.joinRequests = json.users
            })
        },
        async reactToRequest(requestId, response){
            await fetch("http://localhost:8081/responseGroupRequest", {
                method:"POST",
                credentials: "include",
                body: JSON.stringify({
                    groupId: this.$route.params.id,
                    requestId: requestId,
                    response: response
                })
            }).then(r => r.json()).then(json => {
                console.log(json)
            })
        }
    },

}
</script>