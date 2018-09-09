net_cli=./network.sh
domain=infosec.pusan.ac.kr
org1=develop
org2=research
org3=service



function jenga_usage {
    echo "========================================================================="
    echo " jenga"
    echo "-------------------------------------------------------------------------"
    echo " Commands"
    echo " network : bring up and down a network for Hyperledger."
    echo " chaincode : install and deploy chaincode."
    echo ""
    echo " Flags"
    echo " --domain : Domain name($domain)"
    echo " --org1 : The first organization($org1)"
    echo " --org2 : The second organization($org2)"
    echo " --org3 : The third organization($org3)"
    echo ""

}

function jenga_chaincode_usage {
    echo "========================================================================="
    echo " jenga chaincode"
    echo "-------------------------------------------------------------------------"
    echo " Commands"
    echo " deploy : install and deploy chaincode."
    echo " upgrade : install and upgrade a chaincode."
    echo ""
}

function jenga_chaincode_install_usage {
    echo "========================================================================="
    echo " jenga chaincode install"
    echo "-------------------------------------------------------------------------"
    echo " install command was deprecated. Use deploy or upgrade commands, directly."
    echo " Thease commands will install a chaincode and then perform command-specific"
    echo " behavior."
    echo ""
    echo " Deploy example )"
    echo " ./jenga.sh chaincode deploy org pkg/mycc mycc 1.0 common '{\"Args\":[\"init\",\"a\",\"100\",\"b\",\"100\"]}'"
    echo " ./jenga.sh chaincode deploy org chaincode_package registered_name version channel init"
    echo ""
    echo " Upgrade example )"
    echo " ./jenga.sh chaincode upgrade org pkg/mycc mycc 1.1 common '{\"Args\":[\"init\",\"a\",\"100\",\"b\",\"100\"]}'"
    echo " ./jenga.sh chaincode upgrade org chaincode_package registered_name version channel init"
    echo ""

}

function jenga_chaincode_deploy_usage {
    echo "========================================================================="
    echo " jenga chaincode deploy"
    echo "-------------------------------------------------------------------------"
    echo " Example )"
    echo " ./jenga.sh chaincode deploy org pkg/mycc mycc 1.0 common '{\"Args\":[\"init\",\"a\",\"100\",\"b\",\"100\"]}'"
    echo " ./jenga.sh chaincode deploy org chaincode_package registered_name version channel init"
    echo ""
}

function jenga_chaincode_install {
    jenga_chaincode_install_usage
    exit 0
    # o=$1 ## org
    # c=$2 ## chaincode
    # p=$3 ## registered name
    # v=$4 ## version

    # docker-compose --file ledger/docker-compose-$o.yaml run --rm "cli.$o.$domain" bash -c "CORE_PEER_ADDRESS=peer0.$o.$domain:7051 peer chaincode install -n $c -v $v -p $p && CORE_PEER_ADDRESS=peer1.$o.$domain:7051 peer chaincode install -n $c -v $v -p $p"
}

function jenga_chaincode_build_test {
    echo "$1 chaincode will be tested."
    curr=$(pwd)

    cd chaincode/go/$1
    res=`echo $?`
    cd $curr

    if [ "$res" == "0" ]; then
        echo "$1 chaincode was tested successfully."
    else
        echo "Building $1 chaincode was failed."
        exit 1
    fi
}

function jenga_chaincode_deploy {
    if [ "$#" -ne 6 ]; then
        jenga_chaincode_deploy_usage
        exit 0
    fi
    o=$1 ## org
    c=$2 ## chaincode
    p=$3 ## registered name
    v=$4 ## version
    ch=$5
    i=$6 ## {\"Args\":[\"init\",\"a\",\"100\",\"b\",\"100\"]}

    jenga_chaincode_build_test $c

    docker-compose --file ledger/docker-compose-$o.yaml run --rm "cli.$o.$domain" bash -c "CORE_PEER_ADDRESS=peer0.$o.$domain:7051 peer chaincode install -n $c -v $v -p $p && CORE_PEER_ADDRESS=peer1.$o.$domain:7051 peer chaincode install -n $c -v $v -p $p"

    docker-compose --file ledger/docker-compose-$o.yaml run --rm "cli.$o.$domain" bash -c "CORE_PEER_ADDRESS=peer0.$o.$domain:7051 peer chaincode instantiate -n $c -v $v -c '$i' -o orderer.$domain:7050 -C $ch --tls --cafile /etc/hyperledger/crypto/orderer/tls/ca.crt"
}

function jenga_chaincode_upgrade_usage {
    echo "========================================================================="
    echo " jenga chaincode deploy"
    echo "-------------------------------------------------------------------------"
    echo " Example )"
    echo " ./jenga.sh chaincode upgrade org pkg/mycc mycc 1.1 common '{\"Args\":[\"init\",\"a\",\"100\",\"b\",\"100\"]}'"
    echo " ./jenga.sh chaincode upgrade org chaincode_package registered_name version channel init"
    echo ""
}

function jenga_chaincode_upgrade {
    if [ "$#" -ne 6 ]; then
        jenga_chaincode_upgrade_usage
        exit 0
    fi
    o=$1 ## org
    c=$2 ## chaincode
    p=$3 ## registered name
    v=$4 ## version
    ch=$5
    i=$6

    jenga_chaincode_build_test $c

    docker-compose --file ledger/docker-compose-$o.yaml run --rm "cli.$o.$domain" bash -c "CORE_PEER_ADDRESS=peer0.$o.$domain:7051 peer chaincode install -n $c -v $v -p $p && CORE_PEER_ADDRESS=peer1.$o.$domain:7051 peer chaincode install -n $c -v $v -p $p"

    docker-compose --file ledger/docker-compose-$o.yaml run --rm "cli.$o.$domain" bash -c "CORE_PEER_ADDRESS=peer0.$o.$domain:7051 peer chaincode upgrade -n $c -v $v -c '$i' -o orderer.$domain:7050 -C $ch --tls --cafile /etc/hyperledger/crypto/orderer/tls/ca.crt"
}

function jenga_chaincode {
    case $1 in
        install | deploy | upgrade )
            c=jenga_chaincode_$1
            shift
            $c $@
            ;;
        * )
            jenga_chaincode_usage
            ;;
    esac
}

function jenga_network_usage {
    echo "========================================================================="
    echo " jenga network"
    echo "-------------------------------------------------------------------------"
    echo " Commands"
    echo " start : bring up network for Hyperledger."
    echo " stop : bring down network."
    echo " generate : bring down network."
    echo ""
    echo " Flags"
    echo " --domain : Domain name($domain)"
    echo " --org1 : The first organization($org1)"
    echo " --org2 : The second organization($org2)"
    echo " --org3 : The third organization($org3)"
    echo ""

}

function jenga_network_generate {
    $net_cli -m generate
}

function jenga_network_start {
    $net_cli -m up

    ## deploy chaincodes
    for d in $org1
    do
        jenga_chaincode_deploy $d user user 1.0 common '{"Args":["init"]}'
        jenga_chaincode_deploy $d project project 1.0 common '{"Args":["init"]}'
    done
}

function jenga_network_stop {
    $net_cli -m down
}

function jenga_network {
    case $1 in
        start | stop | generate )
            c=jenga_network_$1
            shift
            $c
            ;;
        * )
            jenga_network_usage
            ;;
    esac
}

function jenga_main {
    case $1 in
        chaincode | network )
            c=jenga_$1
            shift
            $c $@
            ;;
        *)
            jenga_usage
            ;;
    esac
}

jenga_main $@
