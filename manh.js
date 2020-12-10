require('./model/mongo-db');
require('./model/Ticket/TicketModel');
const mongoose = require("mongoose");
const TicketModel = mongoose.model("ticketModel");

async function name(params) {
    // await TicketModel.updateMany({
    //     roundId: "274f7ac0",
    //     roi: { $lte: 10.07 - 0.30625 }
    // }, {
    //     $inc: { roi: 0.30625 }
    // })
    let a = await TicketModel.find({
        roundId: "274f7ac0",
        roi: { $lte: 10.07 - 0.30625 }
    })
    // console.log(a);
}


function b(params) {
    name()
}
b()