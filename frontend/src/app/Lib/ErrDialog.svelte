<script>
    import MsgDialog from "./MsgDialog.svelte";

    export let lng = {};
    export const err = {
        show(msg) {
            dialog.show(msg);
        },
    };
    let dialog;
</script>

<MsgDialog
    bind:msg={dialog}
    header={lng.error}
    icon={{ name: "close-circle", color: "var(--red)" }}
    btns={[{ title: lng.close, class: "secondary" }]}
/>
