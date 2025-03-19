<script>
    import { beforeUpdate } from "svelte";
    import BtnIcon from "./BtnIcon.svelte";
    import Icon from "./Icon.svelte";

    export let type = "success";
    export let msg = "";
    export let show = false;
    export let noclose = false;

    let aclass;
    let icon;
    let iconcolor;

    beforeUpdate(() => {
        switch (type) {
            case "success":
                aclass = "alert-success";
                icon = "check-circle";
                iconcolor = "var(--alert-success-icon)";
                break;
            case "warning":
                aclass = "alert-warning";
                icon = "exclamation-circle";
                iconcolor = "var(--alert-warning-icon)";
                break;
        }
    });
</script>

{#if show}
    <div class="alert {aclass}" style="max-width: 40em;">
        <div class="hflex">
            <div class="hflex">
                <Icon name={icon} color={iconcolor} size="2.3em" /><span
                    style="padding-left:0.5em; padding-right: 3em">{msg}</span
                >
            </div>
            {#if !noclose}
                <BtnIcon
                    fn={() => {
                        show = false;
                    }}
                    name="close"
                    size="2em"
                />
            {/if}
        </div>
    </div>
{/if}
