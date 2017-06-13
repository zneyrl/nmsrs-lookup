$(function () {
    var select_all_formal_edu = $("#select_all_formal_edu");
    var delete_formal_edu_button = $("#delete_formal_edu_button");

    $("#year_graduated").datetimepicker({
        viewMode: "years",
        format: "YYYY"
    });

    $("#last_attended").datetimepicker({
        viewMode: "years",
        format: "YYYY/MM"
    });

    $("#formal_edu_form").parsley();
    $("#formal_edu_form").on("submit", function (e) {
        e.preventDefault();
        var high_grade_comp_id_val = $("#high_grade_comp_id").select2("val");
        var course_degree_id_val = $("#course_degree_id").select2("val");
        var school_univ_id_val = $("#school_univ_id").select2("val") != null ? $("#school_univ_id").select2("val") : "";
        var school_univ_other_val = $("#school_univ_other").val().toUpperCase();
        var school_univ_text = school_univ_other_val;
        var year_graduated_val = $("#year_graduated").val();
        var last_attended_val = $("#last_attended").val();

        if (!$("#sunl").prop("checked")) {
            school_univ_text = $("#school_univ_id").select2("data")[0].text;
        }

        switch ($(this).attr("data-action")) {
            case "add":
                var formal_edu_index = 1 + $("#formal_edu_table tbody tr").length++;
                var row = `
                <tr data-index="` + formal_edu_index + `">
                    <td class="formal-edu-checkbox">
                        <input type="checkbox" class="checkbox" id="formal_edu_checkbox_` + formal_edu_index + `">
                    </td>
                    <td class="high-grade-comp">
                        <span>` + $("#high_grade_comp_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="high_grade_comp_id[]" value="` + high_grade_comp_id_val + `">
                    </td>
                    <td class="course-degree">
                        <span>` + $("#course_degree_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="course_degree_id[]" value="` + course_degree_id_val + `">
                    </td>
                    <td class="school-univ">
                        <span>` + school_univ_text + `</span>
                        <input type="hidden" name="school_univ_id[]" value="` + school_univ_id_val + `">
                        <input type="hidden" name="school_univ_other[]" value="` + school_univ_other_val + `">
                    </td>
                    <td class="year-graduated">
                        <span>` + year_graduated_val + `</span>
                        <input type="hidden" name="year_graduated[]" value="` + year_graduated_val + `">
                    </td>
                    <td class="last-attended">
                        <span>` + last_attended_val + `</span>
                        <input type="hidden" name="last_attended[]" value="` + last_attended_val + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="formal-edu-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#formal_edu_table tbody").append(row);
                break;
            case "edit":
                var tr = $("#formal_edu_table tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".high-grade-comp").find("span").text($("#high_grade_comp_id").select2("data")[0].text);
                tr.find(".high-grade-comp").find('input[name="high_grade_comp_id[]"]').val(high_grade_comp_id_val);
                tr.find(".course-degree").find("span").text($("#course_degree_id").select2("data")[0].text);
                tr.find(".course-degree").find('input[name="course_degree_id[]"]').val(course_degree_id_val);
                tr.find(".school-univ").find("span").text(school_univ_text);
                tr.find(".school-univ").find('input[name="school_univ_id[]"]').val(school_univ_id_val);
                tr.find(".school-univ").find('input[name="school_univ_other[]"]').val(school_univ_other_val);
                tr.find(".year-graduated").find("span").text(year_graduated_val);
                tr.find(".year-graduated").find('input[name="year_graduated[]"]').val(year_graduated_val);
                tr.find(".last-attended").find("span").text(last_attended_val);
                tr.find(".last-attended").find('input[name="last_attended[]"]').val(last_attended_val);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".formal-edu-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                select_all_formal_edu.prop("checked", false);
            }

            if ($(".formal-edu-checkbox input:checked").length == $(".formal-edu-checkbox input").length) {
                select_all_formal_edu.prop("checked", true);
            }

            if ($(".formal-edu-checkbox input:checked").length == 0) {
                delete_formal_edu_button.prop("disabled", true);
            } else {
                delete_formal_edu_button.prop("disabled", false);
            }
        });

        $(".formal-edu-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var high_grade_comp_id_val = tr.find(".high-grade-comp").find('input[name="high_grade_comp_id[]"]').val();
            var course_degree_id_val = tr.find(".course-degree").find('input[name="course_degree_id[]"]').val();
            var school_univ_id_val = tr.find(".school-univ").find('input[name="school_univ_id[]"]').val();
            var year_graduated_val = tr.find(".year-graduated").find('input[name="year_graduated[]"]').val();
            var last_attended_val = tr.find(".last-attended").find('input[name="last_attended[]"]').val();

            $("#high_grade_comp_id").val(parseInt(high_grade_comp_id_val)).trigger("change");
            $("#course_degree_id").val(parseInt(course_degree_id_val)).trigger("change");

            if (school_univ_id_val == "") {
                var school_univ_other_val = tr.find(".school-univ").find('input[name="school_univ_other[]"]').val();
                $("#sunl").prop("checked", true).trigger("change");
                $("#school_univ_other").val(school_univ_other_val).trigger("change");
            } else {
                $("#school_univ_id").val(parseInt(school_univ_id_val)).trigger("change");
            }
            $("#year_graduated").val(year_graduated_val).trigger("change");
            $("#last_attended").val(last_attended_val).trigger("change");
            $("#formal_edu_form").attr("data-edit-index", tr.data("index"));
            $("#formal_edu_form").attr("data-action", "edit");
            $("#formal_edu_modal").modal("show");
        });
        select_all_formal_edu.prop("checked", false);
        $("#formal_edu_modal").modal("hide");
    });

    select_all_formal_edu.on("change", function () {
        if ($(".formal-edu-checkbox input").length > 0) {
            $(".formal-edu-checkbox input").prop("checked", $(this).prop("checked"));
            delete_formal_edu_button.prop("disabled", !$(this).prop("checked"));
        }
    });

    delete_formal_edu_button.on("click", function () {
        $(".formal-edu-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                select_all_formal_edu.prop("checked", false);
            }
        });

        if ($("#formal_edu_table tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#formal_edu_modal").on("hidden.bs.modal", function () {
        $("#formal_edu_form").removeAttr("data-action");

        if ($("#school_univ_id").prop("disabled")) {
            $("#school_univ_id").prop("disabled", false);
            $("#school_univ_other").prop("disabled", true);
        }
        $("#high_grade_comp_id").removeAttr("data-parsley-required");
        $("#high_grade_comp_id").val(null).trigger("change");
        $("#high_grade_comp_id").attr("data-parsley-required", true);
        $("#course_degree_id").removeAttr("data-parsley-required");
        $("#course_degree_id").val(null).trigger("change");
        $("#course_degree_id").attr("data-parsley-required", true);
        $("#sunl").prop("checked", false).trigger("change");
        $("#year_graduated").val("");
        $("#last_attended").val("");
    });

    $("#add_formal_edu_button").on("click", function () {
        $("#formal_edu_form").attr("data-action", "add");
        $("#formal_edu_modal").modal("show");
    });

    $("#sunl").on("change", function () {
        $("#school_univ_id").removeAttr("data-parsley-required");
        $("#school_univ_id").val(null).trigger("change");

        if ($(this).prop("checked")) {
            $("#school_univ_id").prop("disabled", true);
            $("#school_univ_other").attr("data-parsley-required", true);
            $("#school_univ_other").prop("disabled", false);
            $("#school_univ_other").focus();
        } else {
            $("#school_univ_other").removeAttr("data-parsley-required");
            $("#school_univ_other").val("");
            $("#school_univ_other").prop("disabled", true);
            $("#school_univ_id").attr("data-parsley-required", true);
            $("#school_univ_id").prop("disabled", false);
            $("#school_univ_id").focus();
        }
    });
    var select_all_pro_license = $("#select_all_pro_license");
    var delete_pro_license_button = $("#delete_pro_license_button");

    $("#pled").datetimepicker({
        viewMode: "years",
        format: "YYYY-MM"
    });

    $("#pro_license_form").parsley();
    $("#pro_license_form").on("submit", function (e) {
        e.preventDefault();
        var plt_id_val = $("#plt_id").select2("val");
        var pled_val = $("#pled").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var pro_license_index = 1 + $("#pro_license_table tbody tr").length++;
                var row = `
                <tr data-index="` + pro_license_index + `">
                    <td class="pro-license-checkbox">
                        <input type="checkbox" class="checkbox" id="pro_license_checkbox_` + pro_license_index + `">
                    </td>
                    <td class="plt">
                        <span>` + $("#plt_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="plt_id[]" value="` + plt_id_val + `">
                    </td>
                    <td class="pled">
                        <span>` + pled_val + `</span>
                        <input type="hidden" name="pled[]" value="` + pled_val + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="pro-license-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#pro_license_table tbody").append(row);
                break;
            case "edit":
                var tr = $("#pro_license_table tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".plt").find("span").text($("#plt_id").select2("data")[0].text);
                tr.find(".plt").find('input[name="plt_id[]"]').val(plt_id_val);
                tr.find(".pled").find("span").text(pled_val);
                tr.find(".pled").find('input[name="pled[]"]').val(pled_val);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".pro-license-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                select_all_pro_license.prop("checked", false);
            }

            if ($(".pro-license-checkbox input:checked").length == $(".pro-license-checkbox input").length) {
                select_all_pro_license.prop("checked", true);
            }

            if ($(".pro-license-checkbox input:checked").length == 0) {
                delete_pro_license_button.prop("disabled", true);
            } else {
                delete_pro_license_button.prop("disabled", false);
            }
        });

        $(".pro-license-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var plt_id_val = tr.find(".plt").find('input[name="plt_id[]"]').val();
            var pled_val = tr.find(".pled").find('input[name="pled[]"]').val();

            $("#plt_id").val(parseInt(plt_id_val)).trigger("change");
            $("#pled").val(pled_val).trigger("change");
            $("#pro_license_form").attr("data-edit-index", tr.data("index"));
            $("#pro_license_form").attr("data-action", "edit");
            $("#pro_license_modal").modal("show");
        });
        select_all_pro_license.prop("checked", false);
        $("#pro_license_modal").modal("hide");
    });

    select_all_pro_license.on("change", function () {
        if ($(".pro-license-checkbox input").length > 0) {
            $(".pro-license-checkbox input").prop("checked", $(this).prop("checked"));
            delete_pro_license_button.prop("disabled", !$(this).prop("checked"));
        }
    });

    delete_pro_license_button.on("click", function () {
        $(".pro-license-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                select_all_pro_license.prop("checked", false);
            }
        });

        if ($("#pro_license_table tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#pro_license_modal").on("hidden.bs.modal", function () {
        $("#pro_license_form").removeAttr("data-action");
        $("#plt_id").removeAttr("data-parsley-required");
        $("#plt_id").val(null).trigger("change");
        $("#plt_id").attr("data-parsley-required", true);
        $("#pled").val("");
        $("#pled").parsley().reset();
    });

    $("#add_pro_license_button").on("click", function () {
        $("#pro_license_form").attr("data-action", "add");
        $("#pro_license_modal").modal("show");
    });
    var select_all_eligibility = $("#select_all_eligibility");
    var delete_eligibility_button = $("#delete_eligibility_button");

    $("#eyt").datetimepicker({
        viewMode: "years",
        format: "YYYY-MM"
    });

    $("#eligibility_form").parsley();
    $("#eligibility_form").on("submit", function (e) {
        e.preventDefault();
        var et_id_val = $("#et_id").select2("val");
        var eyt_val = $("#eyt").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var eligibility_index = 1 + $("#eligibility_table tbody tr").length++;
                var row = `
                <tr data-index="` + eligibility_index + `">
                    <td class="eligibility-checkbox">
                        <input type="checkbox" class="checkbox" id="eligibility_checkbox_` + eligibility_index + `">
                    </td>
                    <td class="et">
                        <span>` + $("#et_id").select2("data")[0].text + `</span>
                        <input type="hidden" name="et_id[]" value="` + et_id_val + `">
                    </td>
                    <td class="eyt">
                        <span>` + eyt_val + `</span>
                        <input type="hidden" name="eyt[]" value="` + eyt_val + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="eligibility-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#eligibility_table tbody").append(row);
                break;
            case "edit":
                var tr = $("#eligibility_table tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".et").find("span").text($("#et_id").select2("data")[0].text);
                tr.find(".et").find('input[name="et_id[]"]').val(et_id_val);
                tr.find(".eyt").find("span").text(eyt_val);
                tr.find(".eyt").find('input[name="eyt[]"]').val(eyt_val);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".eligibility-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                select_all_eligibility.prop("checked", false);
            }

            if ($(".eligibility-checkbox input:checked").length == $(".eligibility-checkbox input").length) {
                select_all_eligibility.prop("checked", true);
            }

            if ($(".eligibility-checkbox input:checked").length == 0) {
                delete_eligibility_button.prop("disabled", true);
            } else {
                delete_eligibility_button.prop("disabled", false);
            }
        });

        $(".eligibility-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var et_id_val = tr.find(".et").find('input[name="et_id[]"]').val();
            var eyt_val = tr.find(".eyt").find('input[name="eyt[]"]').val().toUpperCase();

            $("#et_id").val(parseInt(et_id_val)).trigger("change");
            $("#eyt").val(eyt_val).trigger("change");
            $("#eligibility_form").attr("data-edit-index", tr.data("index"));
            $("#eligibility_form").attr("data-action", "edit");
            $("#eligibility_modal").modal("show");
        });
        select_all_eligibility.prop("checked", false);
        $("#eligibility_modal").modal("hide");
    });

    select_all_eligibility.on("change", function () {
        if ($(".eligibility-checkbox input").length > 0) {
            $(".eligibility-checkbox input").prop("checked", $(this).prop("checked"));
            delete_eligibility_button.prop("disabled", !$(this).prop("checked"));
        }
    });

    delete_eligibility_button.on("click", function () {
        $(".eligibility-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                select_all_eligibility.prop("checked", false);
            }
        });

        if ($("#eligibility_table tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#eligibility_modal").on("hidden.bs.modal", function () {
        $("#eligibility_form").removeAttr("data-action");
        $("#et_id").removeAttr("data-parsley-required");
        $("#et_id").val(null).trigger("change");
        $("#et_id").attr("data-parsley-required", true);
        $("#eyt").val("");
        $("#eyt").parsley().reset();
    });

    $("#add_eligibility_button").on("click", function () {
        $("#eligibility_form").attr("data-action", "add");
        $("#eligibility_modal").modal("show");
    });
    var select_all_vttare = $("#select_all_vttare");
    var delete_vttare_button = $("#delete_vttare_button");

    $("#vttare_form").parsley();
    $("#vttare_form").on("submit", function (e) {
        e.preventDefault();
        var vttare_not_val = $("#vttare_not").val().toUpperCase();
        var vttare_sa_val = $("#vttare_sa").val().toUpperCase();
        var vttare_pote_val = $("#vttare_pote").val().toUpperCase();
        var vttare_cr_val = $("#vttare_cr").val().toUpperCase();
        var vttare_isa_val = $("#vttare_isa").val().toUpperCase();

        switch ($(this).attr("data-action")) {
            case "add":
                var vttare_index = 1 + $("#vttare_table tbody tr").length++;
                var row = `
                <tr data-index="` + vttare_index + `">
                    <td class="vttare-checkbox">
                        <input type="checkbox" class="checkbox" id="vttare_checkbox_` + vttare_index + `">
                    </td>
                    <td class="vttare-not">
                        <span>` + vttare_not_val + `</span>
                        <input type="hidden" name="vttare_not[]" value="` + vttare_not_val + `">
                    </td>
                    <td class="vttare-sa">
                        <span>` + vttare_sa_val + `</span>
                        <input type="hidden" name="vttare_sa[]" value="` + vttare_sa_val + `">
                    </td>
                    <td class="vttare-pote">
                        <span>` + vttare_pote_val + `</span>
                        <input type="hidden" name="vttare_pote[]" value="` + vttare_pote_val + `">
                    </td>
                    <td class="vttare-cr">
                        <span>` + vttare_cr_val + `</span>
                        <input type="hidden" name="vttare_cr[]" value="` + vttare_cr_val + `">
                    </td>
                    <td class="vttare-isa">
                        <span>` + vttare_isa_val + `</span>
                        <input type="hidden" name="vttare_isa[]" value="` + vttare_isa_val + `">
                    </td>
                    <td class="text-center">
                        <a href="#" class="vttare-edit-link"><i class="fa fa-pencil"></i></a>
                    </td>
                </tr>
                `;

                $("#vttare_table tbody").append(row);
                break;
            case "edit":
                var tr = $("#vttare_table tbody").find(`tr[data-index="` + $(this).attr("data-edit-index") + `"]`);
                tr.find(".vttare-not").find("span").text(vttare_not_val);
                tr.find(".vttare-not").find('input[name="vttare_not[]"]').val(vttare_not_val);
                tr.find(".vttare-sa").find("span").text(vttare_sa_val);
                tr.find(".vttare-sa").find('input[name="vttare_sa[]"]').val(vttare_sa_val);
                tr.find(".vttare-pote").find("span").text(vttare_pote_val);
                tr.find(".vttare-pote").find('input[name="vttare_pote[]"]').val(vttare_pote_val);
                tr.find(".vttare-cr").find("span").text(vttare_cr_val);
                tr.find(".vttare-cr").find('input[name="vttare_cr[]"]').val(vttare_cr_val);
                tr.find(".vttare-isa").find("span").text(vttare_isa_val);
                tr.find(".vttare-isa").find('input[name="vttare_isa[]"]').val(vttare_isa_val);
                $(this).removeAttr("data-edit-index");
                break;
        }

        $(".vttare-checkbox input").on("change", function () {
            if ($(this).prop("checked") == false) {
                select_all_vttare.prop("checked", false);
            }

            if ($(".vttare-checkbox input:checked").length == $(".vttare-checkbox input").length) {
                select_all_vttare.prop("checked", true);
            }

            if ($(".vttare-checkbox input:checked").length == 0) {
                delete_vttare_button.prop("disabled", true);
            } else {
                delete_vttare_button.prop("disabled", false);
            }
        });

        $(".vttare-edit-link").on("click", function () {
            var tr = $(this).closest("tr");
            var vttare_not_val = tr.find(".vttare-not").find('input[name="vttare_not[]"]').val();
            var vttare_sa_val = tr.find(".vttare-sa").find('input[name="vttare_sa[]"]').val();
            var vttare_pote_val = tr.find(".vttare-pote").find('input[name="vttare_pote[]"]').val();
            var vttare_cr_val = tr.find(".vttare-cr").find('input[name="vttare_cr[]"]').val();
            var vttare_isa_val = tr.find(".vttare-isa").find('input[name="vttare_isa[]"]').val();

            $("#vttare_not").val(vttare_not_val).trigger("change");
            $("#vttare_sa").val(vttare_sa_val).trigger("change");
            $("#vttare_pote").val(vttare_pote_val).trigger("change");
            $("#vttare_cr").val(vttare_cr_val).trigger("change");
            $("#vttare_isa").val(vttare_isa_val).trigger("change");
            $("#vttare_form").attr("data-edit-index", tr.data("index"));
            $("#vttare_form").attr("data-action", "edit");
            $("#vttare_modal").modal("show");
        });
        select_all_vttare.prop("checked", false);
        $("#vttare_modal").modal("hide");
    });

    select_all_vttare.on("change", function () {
        if ($(".vttare-checkbox input").length > 0) {
            $(".vttare-checkbox input").prop("checked", $(this).prop("checked"));
            delete_vttare_button.prop("disabled", !$(this).prop("checked"));
        }
    });

    delete_vttare_button.on("click", function () {
        $(".vttare-checkbox input").each(function () {
            if ($(this).prop("checked")) {
                $(this).closest("tr").remove();
                select_all_vttare.prop("checked", false);
            }
        });

        if ($("#vttare_table tbody tr").length == 0) {
            $(this).prop("disabled", true);
        }
    });

    $("#vttare_modal").on("hidden.bs.modal", function () {
        $("#vttare_form").removeAttr("data-action");
        $("#vttare_not").val("");
        $("#vttare_not").parsley().reset();
        $("#vttare_sa").val("");
        $("#vttare_sa").parsley().reset();
        $("#vttare_pote").val("");
        $("#vttare_pote").parsley().reset();
        $("#vttare_cr").val("");
        $("#vttare_cr").parsley().reset();
        $("#vttare_isa").val("");
        $("#vttare_isa").parsley().reset();
    });

    $("#add_vttare_button").on("click", function () {
        $("#vttare_form").attr("data-action", "add");
        $("#vttare_modal").modal("show");
    });
});