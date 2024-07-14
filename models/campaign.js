const mongoose = require("mongoose");

const offDaySchema = new mongoose.Schema({
    companyId: {
        type: mongoose.Schema.Types.ObjectId,
        ref: "Company",
    },
    counselorId: {
        type: mongoose.Schema.Types.ObjectId,
        ref: "Counselor",
    },
    date: {
        type: Date,
        required: [true, "Date is required"],
        set: (value) => {
            // Convert the date to UTC
            const date = new Date(value);
            return new Date(
                Date.UTC(date.getFullYear(), date.getMonth(), date.getDate())
            );
        },
    },
    occasion: {
        type: String,
        required: [true, "Occasion field required"],
    },
    offDayCreatedBy: {
        type: mongoose.Schema.Types.ObjectId,
        // required: [true, "Off day created by field required"],
    },
    offDayCreatorRole: {
        type: String,
        // required: [true, "Off day creator role required"],
    },
});

const OffDays = mongoose.model("OffDays", offDaySchema);
module.exports = OffDays;
