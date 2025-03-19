<script>
    import Dialog from "./Dialog.svelte";

    export let header = "";
    export let icon = {};
    export let btns = [];

    export const msg = {
        show(msg) {
            message = msg;
            dialog.show();
        },
    };
    let message;
    let dialog;
</script>

<Dialog bind:mod={dialog} {header} {icon} {btns}>
    <div slot="content" style="padding: 1em 0em">
        <p>{message}</p>
    </div>
</Dialog>
