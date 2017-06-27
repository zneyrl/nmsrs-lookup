$(function () {
    $("#empActivelyLookingForWork").on("change", function () {
        $(this).val($(this).prop("checked"));
    });

    $("#iAccept").on("change", function () {
        $(this).val($(this).prop("checked"));
    });

    $("#personalInfoBirthdate").datetimepicker({
        viewMode: "years",
        format: "YYYY-MM-DD"
    });

    $("#personalInfoBirthdate").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#empPassportNumber").datetimepicker({
        format: "YYYY-MM"
    });

    $("#empPassportNumber").on("dp.change", function() {
        $(this).parsley().validate();
    });

    $("#civilStatRadios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=civilStat]:checked");

        if ($(checked).val() == "OTHER") {
            $("#civilStatOther").attr("data-parsley-required", true);
            $("#civilStatOther").prop("disabled", false);
            $("#civilStatOther").focus();
        } else {
            $("#civilStatOther").removeAttr("data-parsley-required");
            $("#civilStatOther").val("");
            $("#civilStatOther").parsley().reset();
            $("#civilStatOther").prop("disabled", true);
        }
    });

    $("#empStatRadios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=empStat]:checked");

        if ($(checked).val() == "3") {
            $("#empUnEmpStat").val(null).trigger("change");
            $("#empUnEmpStat").attr("data-parsley-required", true);
            $("#empUnEmpStat").prop("disabled", false);
            $("#empUnEmpStat").focus();
        } else {
            $("#empUnEmpStat").removeAttr("data-parsley-required");
            $("#empUnEmpStat").val(null).trigger("change");
            $("#empUnEmpStat").prop("disabled", true);
            $("#empTeminatedOverseasCountry").removeAttr("data-parsley-required");
            $("#empTeminatedOverseasCountry").val(null).trigger("change");
            $("#empTeminatedOverseasCountry").prop("disabled", true);
        }
    });

    $("#disabilityRadios").find("input[type=radio]").on("change", function () {
        var checked = $("input[name=disab]:checked");

        if ($(checked).val() == "OTHER") {
            $("#disabOther").attr("data-parsley-required", true);
            $("#disabOther").prop("disabled", false);
            $("#disabOther").focus();
        } else {
            $("#disabOther").removeAttr("data-parsley-required");
            $("#disabOther").val("");
            $("#disabOther").parsley().reset();
            $("#disabOther").prop("disabled", true);
        }
    });

    $("#disabIsDisabled").on("change", function () {
        $(this).val($(this).prop("checked"));
        $("#disab_0").parsley().reset();

        if ($(this).prop("checked")) {
            $("#disab_0").attr("data-parsley-required", true);
            $("#disabilityRadios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", false);
            });
        } else {
            $("#disab_0").removeAttr("data-parsley-required");
            $("#disabilityRadios").find("input[type=radio]").each(function () {
                $(this).prop("disabled", true);
                $(this).prop("checked", false);
            });
            $("#disabOther").removeAttr("data-parsley-required");
            $("#disabOther").val(null).trigger("change");
            $("#disabOther").prop("disabled", true);
        }
    });

    $("#skillNotListed").on("change", function () {
        $(this).val($(this).prop("checked"));

        if ($(this).prop("checked")) {
            $("#skills").val(null).trigger("change");
            $("#skills").prop("disabled", true);
            $("#skillOther").attr("data-parsley-required", true);
            $("#skillOther").prop("disabled", false);
            $("#skillOther").focus();
        } else {
            $("#skillOther").removeAttr("data-parsley-required");
            $("#skillOther").val("");
            $("#skillOther").parsley().reset();
            $("#skillOther").prop("disabled", true);
            $("#skills").prop("disabled", false);
            $("#skills").focus();
        }
    });

    $("#registeredAt").datetimepicker({
        defaultDate: new Date(),
        format: "YYYY-MM-DD"
    });

    $("#registeredAt").on("dp.change", function() {
        $(this).parsley().validate();
    });

    previewImage("#personalInfoPhoto");

    $("#createRegistrantForm").parsley();
    $("#createRegistrantForm").on("submit", function () {
        duringSubmitDo(this);
    });
});